package time

import (
	"fmt"
	"testing"
	"time"
)

func TestGetUnixTime(t *testing.T) {
	daily := GetDailyTime(time.Now())
	unix := GetUTC8()
	fmt.Println(daily)
	fmt.Println(unix)
	fmt.Println(time.Now().Unix())
	fmt.Println(Format(unix))
}
