package mta

import (
	"errors"
	"strings"
)

// SplitAddress 将 local@domain 拆分为 local 和 domain
func SplitAddress(addr string) (local, domain string, err error) {
	parts := strings.SplitN(addr, "@", 2)
	if len(parts) != 2 {
		return "", "", errors.New("mta: invalid mail address")
	}
	return parts[0], parts[1], nil
}
