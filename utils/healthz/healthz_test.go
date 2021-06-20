package sd

import (
	"fmt"
	"testing"
)

func TestDiskCheck(t *testing.T) {
	disk := DiskCheck()
	fmt.Println(disk)
}

func TestCPUCheck(t *testing.T) {
	cpu := CPUCheck()
	fmt.Println(cpu)
}

func TestRAMCheck(t *testing.T) {
	ram := RAMCheck()
	fmt.Println(ram)
}
