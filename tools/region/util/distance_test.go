package ip2latlong

import (
	"fmt"
	"testing"
)

func TestEarthDistance(t *testing.T) {
	lat1 := 0.0
	long1 := 105.5577
	lat2 := 180.0
	long2 := 105.5577
	distance := EarthDistance(lat1, long1, lat2, long2)
	fmt.Printf("distance:%vm\n ", int64(distance))
	// 	周长 40009880 最大距离为周长的一半
}

// 90.1 ns/op
func BenchmarkEarthDistance(b *testing.B) {
	lat1 := 29.5689
	long1 := 106.5577
	lat2 := 22.5318
	long2 := 114.1374
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EarthDistance(lat1, long1, lat2, long2)
	}
}
