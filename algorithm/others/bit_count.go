package others

import "log"

func main() {
	// 核心为两个位运算：n = n | (n + 1)和n = n & (n - 1)
	// 123->111011,将其标准化是64位,其中的1有6个,其中的0有58个
	log.Println(count0(123))
	log.Println(count1(123))
}

func count0(n int64) int64 {
	var count int64
	for n != -1 { // 退出条件 此时的终止条件应该是每一位都是1也就是1的补码，即-1
		n = n | (n + 1) // 把0转成1
		count++
	}
	return count
}

func count1(n int64) int64 {
	var count int64
	for n != 0 {
		n = n & (n - 1) // 把1转成0
		count++
	}
	return count
}

func count1simple1(n int64) int64 {
	var count int64
	for n != 0 {
		if n%2 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
