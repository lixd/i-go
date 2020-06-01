package main

/*
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
*/
func main() {

}
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for k, v := range row {
		if v {
			for j := 0; j < n; j++ {
				matrix[k][j] = 0
			}
		}
	}

	for k, v := range col {
		if v {
			for i := 0; i < m; i++ {
				matrix[i][k] = 0
			}
		}
	}

}
