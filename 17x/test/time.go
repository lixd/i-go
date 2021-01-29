package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	// t1 := GetUnixTime()
	// fmt.Println(t1)
	// m := make(map[string]int64)
	// m["1"] += 2
	// fmt.Println(m["1"])
	// m["1"] += 2
	// fmt.Println(m["1"])
	// fmt.Println(1608171721 / 30 * 30)
	fmt.Println(time.Now().AddDate(0, 0, 1).Unix())
	unix := time.Now().AddDate(0, 0, 1).Unix()
	year, month, day := time.Unix(unix, 0).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	fmt.Println(daytime.Unix())
}

func GetUnixTime() int64 {
	local := time.FixedZone("UTC", 8*3600)
	year, month, day := time.Now().In(local).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, local).Unix()
	timestamp := daytime
	return timestamp
}

// GetUnixDay 获取到0点的时间戳
func GetUnixDay(unix int64) time.Time {
	year, month, day := time.Unix(unix, 0).Date()
	daytime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return daytime
}

// 模拟一个Redis客户端
type client struct {
	// ... 其他的配置省略
	requestGroup singleflight.Group
}

// 普通查询
func (c *client) Get(key string) (interface{}, error) {
	fmt.Println("Querying Database")
	time.Sleep(time.Second)
	v := "Content of key" + key
	return v, nil
}

// SingleFlight查询
func (c *client) SingleFlightGet(key string) (interface{}, error) {
	v, err, _ := c.requestGroup.Do(key, func() (interface{}, error) {
		return c.Get(key)

	})
	if err != nil {
		return nil, err
	}
	return v, err
}
