package main

import (
	"github.com/sirupsen/logrus"
	_ "i-go/core/conf"
	"i-go/core/db/redisdb"
	"time"

	"fmt"
	"github.com/go-redis/redis"
)

var rc = redisdb.RedisClient

//Redis 增删改查
func main() {
	//RedisString()
	RedisHash()
	//RedisList()
	//RedisSet()
	//RedisZSet()
	//RedisOthers()
}
func RedisHash() {
	// hash
	rc.HSet("illusory", "name", "illusory")
	rc.HSet("illusory", "age", 23)
	rc.HSetNX("illusory", "name", "illusory11")

	get := rc.HGet("illusory", "name")
	logrus.Infof("illusory name:%v", get.Val())
	// HGetAll 返回map结构 直接通过field取值比较方便
	all := rc.HGetAll("illusory")
	logrus.Infof("illusory all:%v", all.Val())
	logrus.Infof("illusory all-name:%v", all.Val()["name"])

	var userMap = map[string]interface{}{
		"name": "illusory",
		"age":  23,
		"addr": "cq",
	}
	rc.HMSet("illusory", userMap)
	// HMGet返回的是数组结构 按照查询的field顺序存储的 只能通过索引取值 field比较多的时候推荐用hgetall
	hmGet := rc.HMGet("illusory", "name", "age", "addr")
	logrus.Infof("illusory:%v", hmGet)

	rc.HDel("illusory", "addr")
	exists := rc.HExists("illusory", "addr")
	logrus.Infof("illusory addr exists:%v", exists)

	rc.HIncrBy("illusory", "age", 2)
	hLen := rc.HLen("illusory")
	logrus.Infof("HLen:%v", hLen)
	keys := rc.HKeys("illusory")
	logrus.Infof("HKeys:%v", keys)
	vals := rc.HVals("illusory")
	logrus.Infof("HVals:%v", vals)

	rc.HScan("illusory", 0, "*", 10)
}

// RedisString  Redis string结构 增删改查
func RedisString() {
	rc.Set("age", 23, 0)
	rc.Set("name", "illusory", 0)

	age := rc.Get("age")
	name := rc.Get("name")
	logrus.Infof("age:%v name:%s", age.Val(), name.Val())

	rc.MSet("age", 23, "name", "illusory")
	mget := rc.MGet("age", "name")
	logrus.Infof("age:%v name:%s", mget.Val()[0], mget.Val()[1])

	rc.Incr("age")
	rc.IncrBy("age", 2)
	rc.Decr("age")
	rc.DecrBy("age", 2)

	rc.Append("name", "newappend")
	rc.Exists("name")
	rc.Expire("name", time.Second)
	time.Sleep(time.Second)
	rc.Exists("name")
	dump := rc.Dump("age")
	logrus.Infof("dump:%v", dump)
}

func RedisOthers() {
	size := rc.DBSize()
	fmt.Printf("rc.DBSize()  %v \n", size)

	info := rc.Info()
	fmt.Printf("rc Info  %v \n", info.Val())
}

func RedisZSet() {
	// ZSet
	fmt.Println("-----------ZSet-------------")
	i17, err := rc.ZAdd("score", redis.Z{22.2, "Java"}, redis.Z{33.5, "Golang"}, redis.Z{44.5, "Python"}, redis.Z{55.5, "JavaScript"}).Result()
	if err != nil {
		fmt.Printf("rc ZAdd error= %v \n", err)
	}
	fmt.Printf("rc ZAdd result= %v \n", i17)
	score := rc.ZScore("score", "Java")
	fmt.Printf("rc ZScore result= %v \n", score)
	i18, err := rc.ZRank("score", "Java").Result()
	fmt.Printf("rc ZAdd result= %v \n", i18)
	i19, err := rc.ZRem("score", "Java").Result()
	fmt.Printf("rc ZAdd result= %v \n", i19)
	fmt.Println("-------------------------")
}

func RedisSet() {
	// set
	fmt.Println("--------------Set-----------------")
	// SAdd 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略
	i10, err := rc.SAdd("job", "Go Modules", "Redis", "MongoDB").Result()
	if err != nil {
		fmt.Printf("rc SAdd error= %v \n", err)
	}
	i11, err := rc.SAdd("todo", "Go Modules", "MongoDB").Result()
	fmt.Printf("rc LPop result= %v \n", i10)
	fmt.Printf("rc LPop result= %v \n", i11)
	// SPop 移除并返回集合中的一个随机元素
	i12, err := rc.SPop("job").Result()
	fmt.Printf("rc LPop result= %v \n", i12)
	i13, err := rc.SInter("job", "todo").Result()
	for _, value := range i13 {
		fmt.Printf("job todo的交集 value=%v \n", value)
	}
	i14, err := rc.SUnion("job", "todo").Result()
	for _, value := range i14 {
		fmt.Printf("job todo的并集 value=%v \n", value)
	}
	i15, err := rc.SDiff("job", "todo").Result()
	for _, value := range i15 {
		fmt.Printf("job todo的差集 value=%v \n", value)
	}
	i16, err := rc.SMembers("job").Result()
	for _, value := range i16 {
		fmt.Printf("job所有元素 value=%v \n", value)
	}
}

func RedisList() {
	//List
	fmt.Println("--------------List-----------------")
	// Lpush left将一个或多个值插入到列表头部
	i6, err := rc.LPush("msg", "l-hello", "l-world", "l-golang", "l-redis").Result()
	if err != nil {
		fmt.Printf("rc LPush error= %v \n", err)
	}
	fmt.Printf("rc LPush result= %v \n", i6)
	// Rpush right 将一个或多个值插入到列表尾部
	i7, err := rc.RPush("msg", "r-hello", "r-world", "r-golang", "r-redis").Result()
	fmt.Printf("rc RPush result= %v \n", i7)
	// LPop 移除并获取列表的第一个元素
	i8, err := rc.LPop("msg").Result()
	fmt.Printf("rc LPop result= %v \n", i8)
	// RPop 移除并获取列表的倒数第一个元素
	i9, err := rc.RPop("msg").Result()
	fmt.Printf("rc LPop result= %v \n", i9)
	// LRange 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定。
	strings, err := rc.LRange("msg", 1, 2).Result()
	fmt.Printf("rc LRange result= %v \n", strings)
}
