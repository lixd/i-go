package main

import (
	"github.com/sirupsen/logrus"
	_ "i-go/core/conf"
	"i-go/core/db/redisdb"
	"time"

	"github.com/go-redis/redis"
)

var rc = redisdb.RedisClient

// Redis 增删改查
func main() {
	// RedisString()
	// RedisHash()
	// RedisList()
	// RedisSet()
	// RedisZSet()
	// RedisOthers()
	// RedisKey()
	RedisHyperLogLog()
}
func RedisHyperLogLog() {
	for i := 0; i < 10; i++ {
		rc.PFAdd("blackNum", i)
	}
	// 判定当前元素是否存在
	// 1.计算count
	// 2.把元素添加进去
	// 3.在计算一次count
	// 4.如果count增加了则说明元素之前是不存在的
	// 10
	pfCountBefore := rc.PFCount("blackNum")
	rc.PFAdd("blackNum", 11)
	// 11
	pfCountAfter := rc.PFCount("blackNum")
	logrus.Infof("before:%v after:%v", pfCountBefore, pfCountAfter)
}

func RedisKey() {
	rc.Del("firstkey")
	rc.Exists("firstKey")
	rc.Expire("firstKey", time.Second)
	rc.TTL("firstKey")
	// 	移除过期时间
	rc.Persist("firstKey")
	// 查询出所有以frist开头的key
	rc.Keys("first*")

	rc.RandomKey()
	// 当secondKey不存在时将firstKey名字修改为secondKey
	rc.RenameNX("firstKey", "secondKey")

	rc.Type("secondKey")

}

func RedisOthers() {
	size := rc.DBSize()
	logrus.Infof("DBSize:%v", size)

	info := rc.Info()
	logrus.Infof("Info:%v", info)
}

func RedisZSet() {

	zAdd := rc.ZAdd("language", redis.Z{22.2, "Java"}, redis.Z{33.5, "Golang"},
		redis.Z{44.5, "Python"}, redis.Z{55.5, "JavaScript"})
	logrus.Infof("ZAdd:%v", zAdd)
	zCount := rc.ZCount("language", "10", "30")
	logrus.Infof("ZCount:%v", zCount)
	rc.ZIncrBy("language", 10, "Java")
	// J开头的成员
	zLexCount := rc.ZLexCount("language", "[J", "(K")
	logrus.Infof("ZLexCount:%v", zLexCount)

	zRange := rc.ZRange("language", 0, 1)
	logrus.Infof("ZRange:%v", zRange)
	/*	// ZRangeByLex必须要所有成员分数相同时才准确
		rc.ZAdd("language", redis.Z{22.2, "Java"}, redis.Z{22.2, "Golang"},
			redis.Z{22.2, "Python"}, redis.Z{22.2, "JavaScript"})
		zRangeByLex := rc.ZRangeByLex("language", redis.ZRangeBy{Min: "[J", Max: "[K"})
		logrus.Infof("ZRangeByLex:%v", zRangeByLex)*/

	zRangeByScore := rc.ZRangeByScore("language", redis.ZRangeBy{Min: "30", Max: "40"})
	logrus.Infof("ZRangeByScore:%v", zRangeByScore)

	zRank := rc.ZRank("language", "Golang")
	logrus.Infof("ZRank:%v", zRank)

	/*	rc.ZRem("language", "Golang")
		// 移除指定得分区间的成员
		rc.ZRemRangeByScore("language", "30", "33")
		// 通过名字移除
		rc.ZRemRangeByLex("language", "[G", "(H")
		// 通过排名索引移除 移除倒数第一名和倒数第二名
		rc.ZRemRangeByRank("language", 0, 1)*/

	// 和不带Rev的相同 只是排序方式换了 默认是得分从小到达 这个是从打大到小
	rc.ZRevRank("language", "Golang")
	rc.ZRevRange("language", 0, 1)
	zRevRangeByLex := rc.ZRevRangeByLex("language", redis.ZRangeBy{Min: "[J", Max: "[K"})
	logrus.Infof("ZRevRangeByLex:%v", zRevRangeByLex)

}

func RedisSet() {
	// SAdd 将一个或多个 member 元素加入到集合 key 当中，已经存在于集合的 member 元素将被忽略
	sAdd := rc.SAdd("golang", "etcd", "gin", "nats")
	logrus.Infof("SAdd:%v", sAdd)
	rc.SAdd("Java", "spring", "mybatis", "tomcat")
	sCard := rc.SCard("golang")
	logrus.Infof("SCard:%v", sCard)

	rc.SRem("golang", "nats")
	// 差集
	sDiff := rc.SDiff("golang", "Java")
	logrus.Infof("SDiff:%v", sDiff)
	rc.SDiffStore("diff", "golang", "Java")

	rc.SAdd("golang", "tomcat")
	// 交集
	rc.SInterStore("inter", "golang", "Java")
	// 并集
	rc.SUnionStore("union", "golang", "Java")
	isMember := rc.SIsMember("inter", "etcd")
	logrus.Infof("SIsMember:%v", isMember)
	sMembers := rc.SMembers("union")
	logrus.Infof("SMembers:%v", sMembers)
	// 随机移除
	sPop := rc.SPop("union")
	logrus.Infof("SPop:%v", sPop)

	sPopN := rc.SPopN("union", 2)
	logrus.Infof("SPopN:%v", sPopN)

	randMember := rc.SRandMember("union")
	logrus.Infof("SRandMember:%v", randMember)

}

func RedisList() {
	// 表不存在则新建
	lPush := rc.LPush("first", "1", "2")
	logrus.Infof("LPush:%v", lPush)
	// 表不存在则不插入
	lPushX := rc.LPushX("first", "3")
	logrus.Infof("LPushX:%v", lPushX)

	rPush := rc.RPush("second", "11", "22")
	logrus.Infof("RPush:%v", rPush)
	rPushX := rc.RPushX("num", "33")
	logrus.Infof("RPushX:%v", rPushX)

	rc.LPop("second")
	rc.RPop("second")

	// rc.BLPop(time.Second, "first", "second")
	// rc.BRPop(time.Second, "first", "second")

	rc.RPopLPush("first", "second")
	rc.BRPopLPush("first", "second", time.Second)

	lIndex := rc.LIndex("first", 0)
	logrus.Infof("LIndex:%v", lIndex)
	rc.LInsert("first", "before", "3", "33")
	rc.LInsert("first", "after", "3", "23")

	lLen := rc.LLen("first")
	logrus.Infof("LLen:%v", lLen)

	lRange := rc.LRange("first", 0, 2)
	logrus.Infof("LRange:%v", lRange)
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
