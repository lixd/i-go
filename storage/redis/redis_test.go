package main

import (
	"i-go/utils"
	"testing"
	"time"
)

func TestConnRedis(t *testing.T) {
	defer utils.InitLog("redis")()
	time.Sleep(time.Second)
}
