package core

import (
	"fmt"
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
	path := viper.GetString("geo")
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

const (
	LangCN = "zh-CN"
	LangEN = "en"
)

func IP2RegionCN(ip string) string {
	region, err := iP2Region(ip, LangCN)
	if err != nil {
		return "内网IP"
	}
	return region
}

func IP2RegionEN(ip string) string {
	region, err := iP2Region(ip, LangEN)
	if err != nil {
		return "Intranet IP"
	}
	return region
}

func iP2Region(ip, lang string) (string, error) {
	ipFormat := net.ParseIP(ip)
	if ipFormat == nil {
		return "", errors.Wrapf(ErrInvalidIP, "IP%s 无法解析", ip)
	}
	city, err := DB.City(ipFormat)
	if err != nil {
		return "", errors.Wrapf(err, "根据IP%s查询City数据", ip)
	}
	printCity(city)
	var spilt string
	if lang == LangEN {
		spilt = " " // 英文时用空格分开
	}
	region := city.Country.Names[lang] + spilt + city.City.Names[lang]
	return region, nil
}

func printCity(city *geoip2.City) {
	fmt.Printf("country: %+v\n", city.Country)
	fmt.Printf("city: %+v\n", city.City)
	fmt.Printf("RegisteredCountry: %+v\n", city.RegisteredCountry)
	fmt.Printf("RepresentedCountry: %+v\n", city.RepresentedCountry)
	fmt.Printf("Location: %+v\n", city.Location)
	fmt.Printf("Continent: %+v\n", city.Continent)
	fmt.Printf("Postal: %+v\n", city.Postal)
	fmt.Printf("Subdivisions: %+v\n", city.Subdivisions)
	fmt.Printf("Traits: %+v\n", city.Traits)
}
