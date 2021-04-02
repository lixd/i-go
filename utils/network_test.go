package utils

import (
	"fmt"
	"testing"
)

func TestGetIntranetIP(t *testing.T) {
	ip := GetIntranetIP()
	fmt.Println(ip)
}
