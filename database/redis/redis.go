package main

import (
	"i-go/core/conf"
	"i-go/core/db/redisdb"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
	"i-go/utils"
	"time"

	"github.com/go-redis/redis"
)

func Init(path string) {
	err := conf.Init(path)
	if err != nil {
		panic(err)
	}
}

// Redis 增删改查
func main() {
	Init("conf/config.yml")
	redisdb.Init()
	//Inc()
	// RedisString()
	// RedisHash()
	//RedisList()
	// RedisSet()
	// RedisZSet()
	// RedisOthers()
	// RedisKey()
	//RedisHyperLogLog()
	RedisPipeline()
	//BloomFilter()
}

const (
	DefaultNum = 10000
)

func Inc() {
	num, err := redisdb.Client().IncrBy("p-questionNum", 1).Result()
	if err != nil {
		logrus.Error(err)
	}
	num += DefaultNum
	logrus.Info(num)
}
func BloomFilter() {
	defer utils.Trace("BloomFilter")()
	var (
		key  = "firstKey"
		data = []byte("bloomFilter")
	)
	bf := NewBloomFilter(1000*20, 3, redisdb.Client())
	bf.Set(key, data)
	isContains := bf.isContains(key, []byte("newData"))
	logrus.Infof("res:%v", isContains)
}

func RedisPipeline() {
	simple()
	pipeline()
}

func pipeline() {
	defer utils.Trace("Redis pipeline写入")()
	pipeLiner := redisdb.Client().Pipeline()

	for i := 0; i < 1000; i++ {
		key := "pipeline" + strconv.Itoa(i)
		pipeLiner.Set(key, i, time.Second*20)
	}
	exec, err := pipeLiner.Exec()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Println("result:", exec)
}

func simple() {
	defer utils.Trace("Redis单条写入")()
	var wg = &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			for j := 0; j < 100; j++ {
				key := "simple" + strconv.Itoa(v) + "g" + strconv.Itoa(j)
				_, err := redisdb.Client().Set(key, j, time.Second*20).Result()
				if err != nil {
					logrus.Error(err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func RedisHyperLogLog() {
	var (
		key    = "clickStatics"
		userId = 10010
	)
	// 删除旧测试数据
	redisdb.Client().Del(key)
	for i := 10000; i < 10010; i++ {
		redisdb.Client().PFAdd(key, i)
	}
	// 判定当前元素是否存在
	// PFAdd添加后对基数值产生影响则返回1 否则返回0
	res := redisdb.Client().PFAdd(key, userId)
	if err := res.Err(); err != nil {
		logrus.Errorf("err :%v", err)
		return
	}
	if res.Val() != 1 {
		logrus.Println("该用户已统计")
	} else {
		logrus.Println("该用户未统计")
	}
}

func RedisKey() {
	redisdb.Client().Del("firstkey")
	redisdb.Client().Exists("firstKey")
	redisdb.Client().Expire("firstKey", time.Second)
	redisdb.Client().TTL("firstKey")
	// 	移除过期时间
	redisdb.Client().Persist("firstKey")
	// 查询出所有以frist开头的key
	redisdb.Client().Keys("first*")

	redisdb.Client().RandomKey()
	// 当secondKey不存在时将firstKey名字修改为secondKey
	redisdb.Client().RenameNX("firstKey", "secondKey")

	redisdb.Client().Type("secondKey")

}

func RedisOthers() {
	size := redisdb.Client().DBSize()
	logrus.Infof("DBSize:%v", size)

	info := redisdb.Client().Info()
	logrus.Infof("Info:%v", info)
}

func RedisZSet() {

	zAdd := redisdb.Client().ZAdd("language", redis.Z{22.2, "Java"}, redis.Z{33.5, "Golang"},
		redis.Z{44.5, "Python"}, redis.Z{55.5, "JavaScript"})
	logrus.Infof("ZAdd:%v", zAdd)
	zCount := redisdb.Client().ZCount("language", "10", "30")
	logrus.Infof("ZCount:%v", zCount)
	redisdb.Client().ZIncrBy("language", 10, "Java")
	// J开头的成员
	zLexCount := redisdb.Client().ZLexCount("language", "[J", "(K")
	logrus.Infof("ZLexCount:%v", zLexCount)

	zRange := redisdb.Client().ZRange("language", 0, 1)
	logrus.Infof("ZRange:%v", zRange)
	/*	// ZRangeByLex必须要所有成员分数相同时才准确
		redisdb.Client().ZAdd("language", redis.Z{22.2, "Java"}, redis.Z{22.2, "Golang"},
			redis.Z{22.2, "Python"}, redis.Z{22.2, "JavaScript"})
		zRangeByLex := redisdb.Client().ZRangeByLex("language", redis.ZRangeBy{Min: "[J", Max: "[K"})
		logrus.Infof("ZRangeByLex:%v", zRangeByLex)*/

	zRangeByScore := redisdb.Client().ZRangeByScore("language", redis.ZRangeBy{Min: "30", Max: "40"})
	logrus.Infof("ZRangeByScore:%v", zRangeByScore)

	zRank := redisdb.Client().ZRank("language", "Golang")
	logrus.Infof("ZRank:%v", zRank)

	/*	redisdb.Client().ZRem("language", "Golang")
		// 移除指定得分区间的成员
		redisdb.Client().ZRemRangeByScore("language", "30", "33")
		// 通过名字移除
		redisdb.Client().ZRemRangeByLex("language", "[G", "(H")
		// 通过排名索引移除 移除倒数第一名和倒数第二名
		redisdb.Client().ZRemRangeByRank("language", 0, 1)*/

	// 和不带Rev的相同 只是排序方式换了 默认是得分从小到达 这个是从打大到小
	redisdb.Client().ZRevRank("language", "Golang")
	redisdb.Client().ZRevRange("language", 0, 1)
	zRevRangeByLex := redisdb.Client().ZRevRangeByLex("language", redis.ZRangeBy{Min: "[J", Max: "[K"})
	logrus.Infof("ZRevRangeByLex:%v", zRevRangeByLex)

}

func RedisSet() {
	// SAdd 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略
	sAdd := redisdb.Client().SAdd("golang", "etcd", "gin", "nats")
	logrus.Infof("SAdd:%v", sAdd)
	redisdb.Client().SAdd("Java", "spring", "mybatis", "tomcat")
	sCard := redisdb.Client().SCard("golang")
	logrus.Infof("SCard:%v", sCard)

	redisdb.Client().SRem("golang", "nats")
	// 差集
	sDiff := redisdb.Client().SDiff("golang", "Java")
	logrus.Infof("SDiff:%v", sDiff)
	redisdb.Client().SDiffStore("diff", "golang", "Java")

	redisdb.Client().SAdd("golang", "tomcat")
	// 交集
	redisdb.Client().SInterStore("inter", "golang", "Java")
	// 并集
	redisdb.Client().SUnionStore("union", "golang", "Java")
	isMember := redisdb.Client().SIsMember("inter", "etcd")
	logrus.Infof("SIsMember:%v", isMember)
	sMembers := redisdb.Client().SMembers("union")
	logrus.Infof("SMembers:%v", sMembers)
	// 随机移除
	sPop := redisdb.Client().SPop("union")
	logrus.Infof("SPop:%v", sPop)

	sPopN := redisdb.Client().SPopN("union", 2)
	logrus.Infof("SPopN:%v", sPopN)

	randMember := redisdb.Client().SRandMember("union")
	logrus.Infof("SRandMember:%v", randMember)

}

func RedisList() {
	// 表不存在则新建
	lPush := redisdb.Client().LPush("first", "1", "2")
	logrus.Infof("LPush:%v", lPush)
	// 表不存在则不插入
	lPushX := redisdb.Client().LPushX("first", "3")
	logrus.Infof("LPushX:%v", lPushX)

	rPush := redisdb.Client().RPush("second", "11", "22")
	logrus.Infof("RPush:%v", rPush)
	rPushX := redisdb.Client().RPushX("num", "33")
	logrus.Infof("RPushX:%v", rPushX)

	redisdb.Client().LPop("second")
	redisdb.Client().RPop("second")

	// redisdb.Client().BLPop(time.Second, "first", "second")
	// redisdb.Client().BRPop(time.Second, "first", "second")

	redisdb.Client().RPopLPush("first", "second")
	redisdb.Client().BRPopLPush("first", "second", time.Second)

	lIndex := redisdb.Client().LIndex("first", 0)
	logrus.Infof("LIndex:%v", lIndex)
	redisdb.Client().LInsert("first", "before", "3", "33")
	redisdb.Client().LInsert("first", "after", "3", "23")

	lLen := redisdb.Client().LLen("first")
	logrus.Infof("LLen:%v", lLen)

	lRange := redisdb.Client().LRange("first", 0, 2)
	logrus.Infof("LRange:%v", lRange)
}

func RedisHash() {
	// hash
	redisdb.Client().HSet("illusory", "name", "illusory")
	redisdb.Client().HSet("illusory", "age", 23)
	redisdb.Client().HSetNX("illusory", "name", "illusory11")

	get := redisdb.Client().HGet("illusory", "name")
	logrus.Infof("illusory name:%v", get.Val())
	// HGetAll 返回map结构 直接通过field取值比较方便
	all := redisdb.Client().HGetAll("illusory")
	logrus.Infof("illusory all:%v", all.Val())
	logrus.Infof("illusory all-name:%v", all.Val()["name"])

	var userMap = map[string]interface{}{
		"name": "illusory",
		"age":  23,
		"addr": "cq",
	}
	redisdb.Client().HMSet("illusory", userMap)
	// HMGet返回的是数组结构 按照查询的field顺序存储的 只能通过索引取值 field比较多的时候推荐用hgetall
	hmGet := redisdb.Client().HMGet("illusory", "name", "age", "addr")
	logrus.Infof("illusory:%v", hmGet)

	redisdb.Client().HDel("illusory", "addr")
	exists := redisdb.Client().HExists("illusory", "addr")
	logrus.Infof("illusory addr exists:%v", exists)

	redisdb.Client().HIncrBy("illusory", "age", 2)
	hLen := redisdb.Client().HLen("illusory")
	logrus.Infof("HLen:%v", hLen)
	keys := redisdb.Client().HKeys("illusory")
	logrus.Infof("HKeys:%v", keys)
	vals := redisdb.Client().HVals("illusory")
	logrus.Infof("HVals:%v", vals)

	redisdb.Client().HScan("illusory", 0, "*", 10)
}

// RedisString  Redis string结构 增删改查
func RedisString() {
	redisdb.Client().Set("age", 23, 0)
	redisdb.Client().Set("name", "illusory", 0)

	age := redisdb.Client().Get("age")
	name := redisdb.Client().Get("name")
	logrus.Infof("age:%v name:%s", age.Val(), name.Val())

	redisdb.Client().MSet("age", 23, "name", "illusory")
	mget := redisdb.Client().MGet("age", "name")
	logrus.Infof("age:%v name:%s", mget.Val()[0], mget.Val()[1])

	redisdb.Client().Incr("age")
	redisdb.Client().IncrBy("age", 2)
	redisdb.Client().Decr("age")
	redisdb.Client().DecrBy("age", 2)

	redisdb.Client().Append("name", "newappend")
	redisdb.Client().Exists("name")
	redisdb.Client().Expire("name", time.Second)
	time.Sleep(time.Second)
	redisdb.Client().Exists("name")
	dump := redisdb.Client().Dump("age")
	logrus.Infof("dump:%v", dump)
}
