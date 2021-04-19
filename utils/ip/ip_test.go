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
