package mta

import (
	"fmt"
	"testing"
)

func TestIDNA(t *testing.T) {
	domainCN := "中文域名.com"
	domainIN, err := ToASCII(domainCN)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("中文域名:%s 国际化域名:%s\n", domainCN, domainIN)
}

func TestIsAllASCII(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "baidu.com"}, want: true},
		{name: "2", args: args{s: "中文域名.com"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllASCII(tt.args.s); got != tt.want {
				t.Errorf("IsAllASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
