package frechet

import (
	"math"
)

/*
参考资料
论文 http://www.kr.tuwien.ac.at/staff/eiter/et-archive/cdtr9464.pdf
http://html.rhhz.net/WHDXXBXXKXB/html/2015-9-1225.htm
https://blog.csdn.net/YYZZHC999/article/details/105799826
*/
type Coordinate struct {
	X int64 `json:"X"`
	Y int64 `json:"Y"`
}
type freChet struct {
	Cord1  []Coordinate
	Cord2  []Coordinate
	Matrix [][]float64
}

var FreChet = &freChet{}

// ClaDiscreteFrechet 计算离散Fréchet距离
func (f *freChet) ClaDiscreteFrechet(one, two []Coordinate) float64 {
	f.Cord1 = one
	f.Cord2 = two
	f.InitFrechetMatrix()
	dfd := f.claDFD(len(one)-1, len(two)-1)
	return dfd
}

// InitFrechetMatrix 初始化矩阵
func (f *freChet) InitFrechetMatrix() {
	var (
		one = f.Cord1
		two = f.Cord2
	)
	matrix := make([][]float64, len(one))
	for i := 0; i < len(one); i++ {
		column := make([]float64, len(two))
		for j := 0; j < len(two); j++ {
			column[j] = -1.0
		}
		matrix[i] = column
	}
	f.Matrix = matrix
}

// claDFD
func (f *freChet) claDFD(i, j int) float64 {
	// if the value has already been solved
	if f.Matrix[i][j] > -1 {
		return f.Matrix[i][j]
	} else if i == 0 && j == 0 {
		// if top left column, just compute the distance
		f.Matrix[i][j] = euclideanDistance(f.Cord1[i], f.Cord2[j])
	} else if i > 0 && j == 0 {
		// can either be the actual distance or distance pulled from above
		f.Matrix[i][j] = max(f.claDFD(i-1, 0), euclideanDistance(f.Cord1[i], f.Cord2[j]))
	} else if i == 0 && j > 0 {
		// can either be the distance pulled from the left or the actual
		// distance
		f.Matrix[i][j] = max(f.claDFD(0, j-1), euclideanDistance(f.Cord1[i], f.Cord2[j]))
	} else if i > 0 && j > 0 {
		// can be the actual distance, or distance from above or from the left
		f.Matrix[i][j] = max(min(f.claDFD(i-1, j), f.claDFD(i-1, j-1), f.claDFD(i, j-1)), euclideanDistance(f.Cord1[i], f.Cord2[j]))
	} else {
		// infinite
		f.Matrix[i][j] = math.MaxInt64
	}
	return f.Matrix[i][j]
}

func min(list ...float64) (min float64) {
	min = math.MaxUint64
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}

func max(list ...float64) (max float64) {
	max = math.MinInt64
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

//  euclideanDistance 欧几里得距离
/*
二维空间: d = sqrt((x1-x2)^2+(y1-y2)^2)
三维空间: d = sqrt((x1-x2)^2+(y1-y2)^2+(z1-z2)^2)
N维空间:  d = sqrt((x1-x2)^2+(y1-y2)^2+(z1-z2)^2+...(n1-n1)^2)
*/
func euclideanDistance(a, b Coordinate) float64 {
	x2 := math.Pow(float64(a.X-b.X), 2)
	y2 := math.Pow(float64(a.Y-b.Y), 2)
	d := math.Sqrt(x2 + y2)
	return d
}
