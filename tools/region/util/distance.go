package ip2latlong

import (
	"math"
)

const (
	radius float64 = 6371000 // 6378137 地球半径(m)
	rad            = math.Pi / 180.0
)

// EarthDistance 球面距离(单位米)
/*
地球是一个近乎标准的椭球体，它的赤道半径为6378.140千米，极半径为 6356.755千米，平均半径6371.004千米（这里忽略地球表面地形对计算带来的误差，仅仅是理论上的估算值）；
设第一点A的经 纬度为(LonA, LatA)，第二点B的经纬度为(LonB, LatB)
按照0度经线的基准，东经取经度的正值(Longitude)，西经取经度负值(-Longitude)，北纬取90-纬度值(90- Latitude)，南纬取90+纬度值(90+Latitude)，
则经过上述处理过后的两点被计为(MLonA, MLatA)和(MLonB, MLatB)。那么根据三角推导，可以得到计算两点距离的如下公式：
C = sin(MLatA)*sin(MLatB)*cos(MLonA-MLonB) + cos(MLatA)*cos(MLatB)
Distance = R*Arccos(C)*Pi/180
*/
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
