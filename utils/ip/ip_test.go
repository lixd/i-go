package ip

import (
	"fmt"
	"testing"
)

const (
	NIP = 3232235777
	AIP = "192.168.1.1"
)

func TestInetNtoA(t *testing.T) {
	ntoA := InetNtoA(NIP)
	fmt.Println(ntoA)
}

func TestInetAtoN(t *testing.T) {
	atoN := InetAtoN(AIP)
	fmt.Println(atoN)
}

func TestExternalIP(t *testing.T) {
	exIP, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exIP)
}

func TestInternalIP(t *testing.T) {
	inIP, err := IntranetIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inIP)
}

func TestInetNtoA1(t *testing.T) {
	type args struct {
		ip int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{ip: NIP}, want: AIP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InetNtoA(tt.args.ip); got != tt.want {
				t.Errorf("InetNtoA() = %v, want %v", got, tt.want)
			}
		})
	}
}
