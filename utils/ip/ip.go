package ip

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
)

// InetNtoA 整形转字符串IP
func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// InetAtoN 字符串IP转整形
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// ExternalIP 获取外网IP
func ExternalIP() (ip string) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	ip = string(ipBytes)
	return
}

// InternalIP 获取内网IP
func InternalIP() (ip string) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}
	return
}
