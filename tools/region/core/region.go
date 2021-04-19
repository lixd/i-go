package core

import (
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/spf13/viper"
)

var (
	client *ip2region.Ip2Region
)

func Init() {
	db := viper.GetString("region")
	if db == "" {
		panic("can not get region")
	}
	var err error
	client, err = ip2region.New(db)
	if err != nil {
		panic(err)
	}
}

func IP2Region(ip string) string {
	ipInfo, err := client.MemorySearch(ip)
	if err != nil {
		return ""
	}
	if ipInfo.Province == ipInfo.City {
		ipInfo.City = ""
	}
	result := ipInfo.Country + ipInfo.Province + ipInfo.City

	return strings.ReplaceAll(result, "0", "")
}
