package frechet

import (
	"fmt"
	"testing"
)

func Test_freChet_ClaDiscreteFrechet(t *testing.T) {
	cord1 := []Coordinate{
		{273, 144}, {267, 147}, {261, 149}, {254, 151}, {248, 152},
		{242, 155}, {233, 156}, {225, 156}, {216, 158}, {211, 159},
	}
	cord2 := []Coordinate{
		{280, 144}, {267, 147}, {261, 149}, {254, 151}, {248, 152},
		{242, 155}, {233, 156}, {225, 156}, {216, 158}, {211, 159},
	}

	frechet := FreChet.ClaDiscreteFrechet(cord1, cord2)
	// frechet := FreChet.ClaDiscreteFrechet(reverse(cord1), cord2)
	fmt.Println("frechet: ", frechet)
	// 具体FreChetDistance与相似度关系需要根据需求调整
	// 例如轨迹最大宽高为400 相似度最大100 那么400=>100 大概相差1降低0.25相似度
}

func reverse(points []Coordinate) []Coordinate {
	for i, j := 0, len(points)-1; i < j; i, j = i+1, j-1 {
		(points)[i], (points)[j] = (points)[j], (points)[i]
	}
	return points
}
