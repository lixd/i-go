package cpucache

const (
	row = 1024
	col = 512
)

/*
逐行遍历可以用到 CPU 缓存,数组在内存中是以行为单位存储的，例如：row1，row2，row3 这样。
CPU 每次从内存中取一行加载到CPU缓存(CPU Cache Line)中，如果按照行读取则可以用到CPU缓存，相比逐列遍历，效率自然更高
*/
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
