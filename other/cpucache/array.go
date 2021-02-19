package cpucache

const (
	row = 1024
	col = 512
)

var (
	matrix = [row][col]int64{}
	sum    int64
)

// Row 逐行遍历
func Row() {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum += matrix[i][j]
		}
	}
}

// Col 逐列遍历
func Col() {
	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			sum += matrix[i][j]
		}
	}
}
