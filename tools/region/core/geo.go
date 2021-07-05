package core

import (
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

/*
根据IP查询经纬度 主要依赖 maxmind 提供的数据库 xxx.mmdb
blog: https://blog.csdn.net/qq_26373925/article/details/111876765
注册地址: https://www.maxmind.com/en/geolite2/signup
下载链接: https://www.maxmind.com/en/accounts/576594/geoip/downloads
*/

var ErrInvalidIP = errors.New("invalid ip")

type LatLong struct {
	IP        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`  // 维度
	Longitude float64 `json:"longitude"` // 经度
}

type GEO interface {
	Query(ip string) (LatLong, error)
}

type geo struct {
	db *geoip2.Reader
}

func NewGEO(db *geoip2.Reader) GEO {
	return &geo{db: db}
}

func (g *geo) Query(ip string) (LatLong, error) {
	var item LatLong
	ipFormat := net.ParseIP(ip)
	if ipFormat == nil {
		return item, errors.Wrap(ErrInvalidIP, "解析IP")
	}
	city, err := g.db.City(ipFormat)
	if err != nil {
		return item, errors.Wrap(err, "查询具体位置")
	}
	item = LatLong{
		Latitude:  city.Location.Latitude,
		Longitude: city.Location.Longitude,
	}
	return item, nil
}

// const path = "D:\\lillusory\\projects\\i-go\\tools\\ip2loc\\GeoLiteCity\\GeoLite2-City.mmdb"

var DB *geoip2.Reader

func InitLatLong() {
	var err error
	path := viper.GetString("latlong")
	if path == "" {
		panic("can not get latlong")
	}
	DB, err = geoip2.Open(path)
	if err != nil {
		panic(err)
	}
}

func IP2LatLong(ip string) (LatLong, error) {
	var item LatLong
	ipFormat := net.ParseIP(ip)
	if ipFormat == nil {
		return item, errors.Wrap(ErrInvalidIP, "解析IP")
	}
	city, err := DB.City(ipFormat)
	if err != nil {
		return item, errors.Wrap(err, "查询具体位置")
	}
	item = LatLong{
		IP:        ip,
		Latitude:  city.Location.Latitude,
		Longitude: city.Location.Longitude,
	}
	return item, nil
}
