package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type config struct {
	SSH     *SSHConfig      `json:"ssh"`
	Tunnels []*TunnelConfig `json:"tunnels"`
}

func main() {
	cfg := &config{
		SSH: &SSHConfig{
			Host:           "172.20.149.53",
			User:           "root",
			Port:           22,
			Secret:         "root",
			PrivateKeyFile: "",
		},
		Tunnels: []*TunnelConfig{
			{
				Ident:      "hello",
				LocalPort:  50051,
				RemoteHost: "192.168.10.193",
				RemotePort: 50051,
			},
		},
	}

	// start tunnels with cfg
	startTunnels(cfg)
}

// 默认规则，匹配全部
var defaultPattern = ".*"
var tunnelIdentPattern = &defaultPattern

// upport tunnel open by ident matcher
func startTunnels(cfg *config) {
	var (
		errChan    = make(chan errTunnel, 1) // 异常channel
		wg         = sync.WaitGroup{}        // 同步组
		tunnelChan = make(chan int, 1)       // 运行中tunnels 计数

		exp *regexp.Regexp // pattern regexp
	)

	// compile pattern
	exp = regexp.MustCompile(*tunnelIdentPattern)

	wg.Add(len(cfg.Tunnels))
	// create and ssh tunnel and goto work
	for idx, v := range cfg.Tunnels {
		if v.SSH == nil {
			v.SSH = cfg.SSH
		}

		// matche with ident pattern
		if !exp.MatchString(v.Ident) {
			log.Printf("tunnel ident=%s, not matched with pattern=%s, so skipped", v.Ident, *tunnelIdentPattern)
			continue
		}

		go func(idx int, tunnelCfg *TunnelConfig, errChan chan<- errTunnel) {
			defer wg.Done()
			defer func() { tunnelChan <- -1 }()
			tunnelChan <- 1

			// open tunnel and prepare
			tunnel := NewSSHTunnel(tunnelCfg)
			if err := tunnel.Start(); err != nil {
				errChan <- newErrTunnel(idx, "tunnel broken, err=%v", err)
				return
			}
		}(idx, v, errChan)
	}

	// log errors
	go func() {
		for err := range errChan {
			log.Printf("tunnelIdx=%d: %s", err.Idx, err.Errmsg)
		}
	}()

	// record changes of opening-tunnel count
	go func() {
		running := 0
		msg := ""
		for cntChange := range tunnelChan {
			// if runningTunnelsCnt changed to notify
			// FIXME: atomic op with running
			running += cntChange
			if cntChange >= 0 {
				// true: starting
				msg = fmt.Sprintf("%d tunnel starting, current: %d", cntChange, running)
			} else {
				// true: quit
				msg = fmt.Sprintf("%d tunnel break, current: %d", 0-cntChange, running)
			}
			log.Printf(msg)
		}
	}()

	wg.Wait()
	close(errChan)
	close(tunnelChan)
	// wait for all error message outputing
	time.Sleep(100 * time.Millisecond)
	log.Printf("tunnel-helper is finished")
}

// errTunnel .
type errTunnel struct {
	Idx    int
	Errmsg string
}

func newErrTunnel(idx int, format string, args ...interface{}) errTunnel {
	return errTunnel{
		Idx:    idx,
		Errmsg: fmt.Sprintf(format, args...),
	}
}

func (err errTunnel) Error() string {
	return err.Errmsg
}

// SSHConfig .
type SSHConfig struct {
	Host           string `json:"host"`
	User           string `json:"user"`
	Port           int    `json:"port"`
	Secret         string `json:"secret"`
	PrivateKeyFile string `json:"privateKeyFile"`
}

// TunnelConfig .
type TunnelConfig struct {
	Ident      string     `json:"ident"`
	SSH        *SSHConfig `json:"ssh"`
	LocalPort  int        `json:"localPort"`
	RemoteHost string     `json:"remoteHost"`
	RemotePort int        `json:"remotePort"`
}

// SSHTunnel .
type SSHTunnel struct {
	LocalAddr  string            // format: "host:port"
	ServerAddr string            // format: "host:port"
	RemoteAddr string            // format: "host:port"
	SSHConfig  *ssh.ClientConfig // ssh client config
}

// output: tunnel=(localhost:6379)
func (tunnel *SSHTunnel) name() string {
	return "tunnel=(" + tunnel.LocalAddr + ")"
}

// NewSSHTunnel .
func NewSSHTunnel(tunnelConfig *TunnelConfig) *SSHTunnel {
	var (
		auth      ssh.AuthMethod
		sshConfig = tunnelConfig.SSH
	)

	if sshConfig.PrivateKeyFile != "" {
		// true: privateKey specified
		auth = loadPrivateKeyFile(sshConfig.PrivateKeyFile)
	} else {
		auth = ssh.Password(sshConfig.Secret)
	}

	return &SSHTunnel{
		SSHConfig: &ssh.ClientConfig{
			User: sshConfig.User,
			Auth: []ssh.AuthMethod{auth},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				// log.Printf("accept hostkey callback")
				//  Always accept key.
				return nil
			},
		},
		LocalAddr:  assembleAddr("localhost", tunnelConfig.LocalPort),
		ServerAddr: assembleAddr(tunnelConfig.SSH.Host, tunnelConfig.SSH.Port),
		RemoteAddr: assembleAddr(tunnelConfig.RemoteHost, tunnelConfig.RemotePort),
	}
}

// format like "host:port"
func assembleAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// Start .
// TODO: support random port by using localhost:0
func (tunnel *SSHTunnel) Start() error {
	listener, err := net.Listen("tcp", tunnel.LocalAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	// tunnel.Local.Port = listener.Addr().(*net.TCPAddr).Port
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		log.Printf(tunnel.name() + " accepted connection")
		go tunnel.forward(conn)
	}
}

// just do the work like proxy to transfer data from local to remote
func (tunnel *SSHTunnel) forward(localConn net.Conn) {
	serverSSHClient, err := ssh.Dial("tcp", tunnel.ServerAddr, tunnel.SSHConfig)
	if err != nil {
		log.Printf(tunnel.name()+" server dial error: %s", err)
		return
	}
	log.Printf(tunnel.name()+" connected to server=%s (1 of 2)", tunnel.ServerAddr)
	remoteConn, err := serverSSHClient.Dial("tcp", tunnel.RemoteAddr)
	if err != nil {
		log.Printf(tunnel.name()+" remote dial error: %s", err)
		return
	}
	log.Printf(tunnel.name()+" connected to remote=%s (2 of 2)", tunnel.RemoteAddr)

	copyConn := func(writer, reader net.Conn) {
		_, err = io.Copy(writer, reader)
		if err != nil {
			log.Printf(tunnel.name()+" io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

// loadPrivateKeyFile . load privare file by @dir
func loadPrivateKeyFile(dir string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
