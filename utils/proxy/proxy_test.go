package proxyutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// func TestMain(m *testing.M) {
// 	var file string
// 	flag.StringVar(&file, "f", "proxyConf/config_user.yaml", "the config file path")
// 	flag.Parse()
//
// 	if err := conf.Init(file); err != nil {
// 		panic(err)
// 	}
// 	logger.Init()
// 	redisdb.Init()
// 	m.Run()
// }

func TestLoadProxyIP(t *testing.T) {
	data, err := LoadPSWithoutCache(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

const proxyServer = "http://114.238.101.164:3617"

func TestNormalProxy(t *testing.T) {
	normalProxy(proxyServer)
}

func TestIsValidProxy(t *testing.T) {
	isValidProxy := IsValidProxy(proxyServer)
	fmt.Println(isValidProxy)
}

func TestBuildClientByProxy(t *testing.T) {
	client, err := BuildClientWithProxy("")
	if err != nil {
		fmt.Println(err)
		return
	}
	request, err := http.NewRequest(http.MethodGet, "http://vaptcha.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

// RemoteIP
// XForwardFor

func Test_socket5Proxy(t *testing.T) {
	socket5Proxy("http://myip.ipip.net", "180.113.10.58:14865")
}

func TestIsValidProxySocks(t *testing.T) {
	socks := IsValidProxySocks("180.113.8.6:49201")
	fmt.Println("result:", socks)
}

func Test_loadProxyIP(t *testing.T) {
	Validate("htt://183.92.73.178:5412")
	return
	ip, err := loadProxyIP(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("proxy:%+v\n", ip)
}
