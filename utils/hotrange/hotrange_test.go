package hotrange

import (
	"fmt"
	"testing"
)

func TestNewtonsLawOfCooling(t *testing.T) {
	cooling := NewtonsLawOfCooling(1, 24)
	fmt.Println(cooling)
}
