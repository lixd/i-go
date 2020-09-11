package main

import (
	"fmt"
)

func main() {
	// newInt := big.NewInt(0)
	// newInt.SetBytes(net.ParseIP("").To4())
	n, i2, ok := dtoi("192.168.1.1")
	fmt.Printf("n%v i2:%v ok:%v \n", n, i2, ok)
}

const big2 = 0xFFFFFF

func dtoi(s string) (n int, i int, ok bool) {
	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big2 {
			return big2, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}
