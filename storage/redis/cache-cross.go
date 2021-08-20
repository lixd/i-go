package main

/*// get 伪代码 互斥锁 防止缓存击穿问题 建议使用 singleflight 库
func get(key string) {
	ret := redis.get(Key)
	if ret == nil { // 为空代表缓存值过期
		// 获取锁 同时设置3min的超时，防止del操作失败的时候，下次缓存过期一直不能load db
		setMutex := redis.setnx(key_mutex, 1, 3*60) == 1
		if setMutex {
			// 成功获取锁则load db并回设到缓存
			value = db.get(key)
			redis.set(key, value, expire_secs)
			redis.del(key_mutex)
		} else {
			// 获取失败表示其他请求已经在load db并回设到缓存了 sleep一会然后重试
			time.Sleep(50)
			get(key)
		}
	}
}
*/
