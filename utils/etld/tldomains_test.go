package etld

import "testing"

func TestTLDomains(t *testing.T) {
	tests := []struct {
		in  string
		out Host
	}{
		{in: "mMmm.jello.co.uk", out: Host{"mmmm", "jello", "co.uk"}},
		{in: "pressly.com", out: Host{"", "pressly", "com"}},
		{in: "www.pressly.it", out: Host{"www", "pressly", "it"}},
	}

	for _, tt := range tests {
		h := Parse(tt.in)
		if h.Subdomain != tt.out.Subdomain || h.Domain != tt.out.Domain || h.Suffix != tt.out.Suffix {
			t.Errorf("expected %v, got %v", tt.out, h)
		}
	}
}
