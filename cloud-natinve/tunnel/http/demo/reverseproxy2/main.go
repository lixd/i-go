package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"i-go/cloud-natinve/tunnel/http/demo"
	"k8s.io/klog/v2"
)

const defaultHost = "127.0.0.1"

var defaultSSHConn ssh.Conn
var defaultNetConn net.Conn

func main() {
	http.HandleFunc("/conn", conn)
	http.HandleFunc("/proxy", proxy)
	err := http.ListenAndServe(":50051", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func conn(w http.ResponseWriter, r *http.Request) {
	klog.V(4).Info("New agent connection")
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		klog.Error("Failed to upgrade connection", err)
		return
	}
	connection := demo.NewWebSocketConn(wsConn)
	defaultNetConn = connection
	klog.V(4).Info("Handshaking...")
	sshConfig := &ssh.ServerConfig{
		ServerVersion:    "SSH-" + "kubesphere-v1" + "-server",
		PasswordCallback: authenticate,
	}
	key, _ := generateKey()
	private, err := ssh.ParsePrivateKey(key)
	if err != nil {
		klog.Fatalf("Failed to parse key %v", err)
	}
	sshConfig.AddHostKey(private)
	fmt.Printf("[reverse proxy server] New Conn at: %s\n", time.Now())

	sshConn, chans, reqs, err := ssh.NewServerConn(connection, sshConfig)
	if err != nil {
		klog.Error("Failed to handshake", err)
		return
	}
	_ = chans
	_ = reqs
	defaultSSHConn = sshConn
	fmt.Printf("[reverse proxy server] Conn Successofuly at: %s\n", time.Now())
	fmt.Printf("[reverse proxy server] localAddr:%s remoteAddr:%s \n", defaultNetConn.LocalAddr(), defaultNetConn.RemoteAddr())
}

func generateKey() ([]byte, error) {
	r := rand.Reader

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), r)
	if err != nil {
		return nil, err
	}
	b, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal ECDSA private key: %v", err)
	}
	return pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b}), nil
}

func proxy(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())
	// define origin server URL
	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}
	// set req Host, URL and Request URI to forward a request to the origin server
	req.Host = originServerURL.Host
	req.URL.Host = originServerURL.Host
	req.URL.Scheme = originServerURL.Scheme
	req.RequestURI = ""
	fmt.Printf("[reverse proxy server] req host:%s url.host:%s url.Scheme:%s uri:%s \n",
		req.Host, req.URL.Host, req.URL.Scheme, req.RequestURI)

	//
	// send a request to the origin server
	// if reverseProxy server can access target server just do it.
	// originResp, err := http.DefaultClient.Do(req)
	// if reverseProxy server can not access target server,deploy proxy agent on target server and send a req to
	// reverseProxy server ,then save this conn and warp to a *http.client,now use this client to transform req.
	cli := getClient()
	originResp, err := cli.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(rw, err)
		return
	}
	// return response to the client
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, originResp.Body)
}

func getClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
				return demo.NewSshConn(func() ssh.Conn {
					return defaultSSHConn
				}, defaultHost), nil
			},
			// DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, err error) {
			// 	return defaultNetConn, nil
			// },
		},
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func authenticate(c ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	klog.V(4).Infof("%s is connecting from %s", c.User(), c.RemoteAddr())
	return nil, nil
}
