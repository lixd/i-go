package middle

/*
面试题 01.07. 旋转矩阵
*/
func main() {
	//var matrix [][]int{[]int{1,2,3},[]{4,5,6},[]int{7,8,9}}
	//rotate(matrix)
	//fmt.Println(matrix)
}
func rotate(matrix [][]int) {

	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
