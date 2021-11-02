package main

import (
	"fmt"

	"i-go/core/conf"
	"i-go/core/db/redisdb"
)

//  直接加载官方实现的 BloomFilter 模块
/*
相关参考链接如下
https://redis.uptrace.dev/guide/bloom-cuckoo-count-min-top-k.html#bloom-vs-cuckoo
https://redis.com/redis-best-practices/bloom-filter-pattern/
https://www.cnblogs.com/wy123/p/11571215.html
加载方法
下载BloomFilter 模块并编译成.so文件,然后在Redis.conf 中配置加载该so文件，最后重启即可
*/

func init() {
	err := conf.Load("conf/config.yml")
	if err != nil {
		panic(err)
	}
	redisdb.Init()
}

func main() {
	rdb := redisdb.Client()
	// 这个驱动暂时没实现 BF 相关命令，只能借助 Do 来执行了
	// 1.先指定允许的错误率和容量
	err := rdb.Do("BF.RESERVE", "bf_key", 0.0001, 10_0000_0000).Err()
	if err != nil {
		fmt.Println("RESERVE err:", err)
		return
	}
	// 2.然后添加数据
	inserted, err := rdb.Do("BF.ADD", "bf_key", "item0").Bool()
	if err != nil {
		panic(err)
	}
	if inserted {
		fmt.Println("item0 was inserted")
	}
	// 3.查看数据是否存在
	for _, item := range []string{"item0", "item1"} {
		exists, err := rdb.Do("BF.EXISTS", "bf_key", item).Bool()
		if err != nil {
			panic(err)
		}
		if exists {
			fmt.Printf("%s does exist\n", item)
		} else {
			fmt.Printf("%s does not exist\n", item)
		}
	}
}
