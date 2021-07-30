package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/*
	level ms
	3 	  1
	4    13 33 53
	5	 192 574
	6	 3282
	7	 75855
6 降低难度(01234) 1153
6 降低难度(012345) 953
*/
func main() {
	powParallel(4, 200)
}

func powParallel(level, limit int) {
	var (
		start        = time.Now().UnixNano()
		globalStart  = time.Now().UnixNano()
		index, count int64
		wg           sync.WaitGroup
	)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				target := int(atomic.LoadInt64(&index))
				// str := "vid" + strconv.Itoa(target)
				str := "13424so7khmbd6d1" + "https://img-cn.vaptcha.net/vaptcha/3a18c9f8c35c49eebd1f29a0e874572e.jpg" + strconv.Itoa(target)
				s := hashSha256([]byte(str))
				if s[:level] == strings.Repeat("0", level) {
					// if isOk(s,"012345",level) {
					atomic.AddInt64(&count, 1)
					fmt.Printf("time:%vms hash: %s, count:%v;\n", (time.Now().UnixNano()-start)/1e6, s, target)
					start = time.Now().UnixNano()
				}
				if atomic.LoadInt64(&count) > int64(limit) {
					break
				}
				atomic.AddInt64(&index, 1)
			}
		}()
	}
	wg.Wait()
	avgTime := (time.Now().UnixNano() - globalStart) / 1e6 / atomic.LoadInt64(&count)
	fmt.Printf("level:%v,avgTime:%v;\n", level, avgTime)
}

func isOk(str, com string, level int) bool {
	for _, v := range com {
		if str[:level+1] == strings.Repeat("0", level)+string(v) {
			return true
		}
	}
	return false
}

func hashSha256(src []byte) string {
	hash := sha256.New()
	// 输入数据
	hash.Write(src)
	// 计算哈希值
	bytes := hash.Sum(nil)
	// 将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	// 返回哈希值
	return hashCode
}
