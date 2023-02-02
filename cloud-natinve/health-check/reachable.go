package health_check

import (
	"context"
	"fmt"
	"github.com/go-ping/ping"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	PacketLossRateLimit = 30
)

// ReachableByPing check reachable by ping
func ReachableByPing(addr string, timeout time.Duration) error {
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		return err
	}
	pinger.Timeout = timeout
	pinger.Count = 10
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return err
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	if stats.PacketLoss >= PacketLossRateLimit {
		return fmt.Errorf("ping [%s] failed,packet loss rate [%.f%%] gather than limit %v%%", addr, stats.PacketLoss, PacketLossRateLimit)
	}
	return nil
}

// ReachableByTCP check reachable by tcp
func ReachableByTCP(host, port string, timeout time.Duration) error {
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	defer connection.Close()
	return nil
}

// ReachableByHTTP check reachable by http(s)
func ReachableByHTTP(protocol, host, port string, timeout time.Duration) error {
	if port == "" {
		if protocol == "http" {
			port = "80"
		} else if protocol == "https" {
			port = "443"
		}
	}
	url := fmt.Sprintf("%s://%s:%s", protocol, host, port)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	if resp.StatusCode >= http.StatusInternalServerError {
		return fmt.Errorf("resp http code %v", resp.StatusCode)
	}
	return nil
}
