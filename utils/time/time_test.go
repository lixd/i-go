package time

import (
	"fmt"
	"testing"
	"time"
)

func TestGetUnixTime(t *testing.T) {
	date := DateTime()
	unix := UnixTime()
	fmt.Println(date)
	fmt.Println(unix)
	fmt.Println(time.Now().Unix())
	fmt.Println(Format(unix))
}
