package core

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

const (
	LocalIP = "49.76.162.72"
)

func TestMain(m *testing.M) {
	viper.Set("geo", "D:\\wlinno\\projects\\vaptcha-go\\conf\\region\\GeoLite2-City.mmdb")
	InitLatLong()
	m.Run()
}

func TestIP2LatLong(t *testing.T) {
	long, err := IP2LatLong(LocalIP)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(long)
}

func TestIP2RegionCN(t *testing.T) {
	cn := IP2RegionCN(LocalIP)
	en := IP2RegionEN(LocalIP)
	fmt.Println(cn, en)
}
