package hotrank

import (
	"fmt"
	"testing"
)

func TestNewtonsLawOfCooling(t *testing.T) {
	cooling := NewtonsLawOfCooling(100, 24)
	fmt.Println(cooling)
}
