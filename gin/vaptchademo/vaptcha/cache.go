package vaptcha

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// ICache 提供存储功能 用于存储 knock
type ICache interface {
	Get(key interface{}) string
	Set(key, value interface{})
	Delete(key interface{})
	Range(f func(key, value interface{}) bool)
}

var (
	mem sync.Map
)

// vCache default cache
type vCache struct {
}

func (vCache) Get(key interface{}) string {
	if value, ok := mem.Load(key); ok {
		return value.(string)
	}
	return ""
}
func (vCache) Set(key, value interface{}) {
	mem.Store(key, value)
}
func (vCache) Delete(key interface{}) {
	mem.Delete(key)
}
func (vCache) Range(f func(key, value interface{}) bool) {
	mem.Range(f)
}

// RemoveExpireSess
func (v *vaptcha) RemoveExpireKey() {
	for {
		f := func(key, value interface{}) bool {
			var val = fmt.Sprintf("%v", value)
			if len(val) > 10 {
				unix, _ := strconv.ParseInt(val[:10], 10, 64)
				timeSpan := time.Now().Unix() - unix
				// 大于3分钟过期
				if timeSpan > 3*60 || timeSpan < -60 {
					v.options.Cache.Delete(key)
				}
			} else {
				// 错误key直接删除
				v.options.Cache.Delete(key)
			}
			return true
		}
		v.options.Cache.Range(f)
		time.Sleep(time.Second * 1)
	}
}
