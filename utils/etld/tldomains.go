package etld

import (
	"fmt"
	"strings"
)

var TLDs = make(map[string]struct{}, 0)

type Host struct {
	Subdomain, Domain, Suffix string
}

func Parse(host string) Host {
	var h Host

	nhost := strings.ToLower(host)
	parts := strings.Split(nhost, ".")

	if len(parts) == 0 {
		h.Domain = host
		return h
	}

	var suffix string
	for i := len(parts) - 1; i >= 0; i-- {
		p := parts[i]

		if suffix == "" {
			suffix = p
		} else {
			suffix = fmt.Sprintf("%s.%s", p, suffix)
		}

		if _, ok := TLDs[suffix]; ok {
			h.Suffix = suffix
		} else if h.Domain == "" {
			h.Domain = p
		} else {
			h.Subdomain = p
		}
	}

	return h
}
