package daily

import (
	"fmt"
	"testing"
)

func Test_monotoneIncreasingDigits(t *testing.T) {
	digits := monotoneIncreasingDigits(1243)
	fmt.Println(digits)
}
