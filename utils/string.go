package utils

import (
	"math"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"i-go/utils/murmur"
)

type stringHelper struct {
}

// StringHelper string相关工具函数
var StringHelper = &stringHelper{}

// GetUUID 生成UUID
/*
几个uuid库
https://github.com/google/uuid
https://github.com/gofrs/uuid
// google 和 gofrs 二者未对比
https://github.com/satori/go.uuid 不推荐 https://github.com/satori/go.uuid/issues/84
*/
func (stringHelper) GetUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

// ShuffleSlice 数组乱序
func ShuffleSlice(arr []string) []string {
	for i := len(arr) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		arr[i], arr[num] = arr[num], arr[i]
	}
	return arr
}

// Subset 从服务列表中取一部分出来进行健康检测
func Subset(backends []string, clientID string, subsetSize int64) []string {
	subSetCount := int64(len(backends)) / subsetSize
	clientHash := hashDemo(clientID)
	round := clientHash / subSetCount
	rand.Seed(round)
	backends = ShuffleSlice(backends)
	subSetID := clientHash % subSetCount
	start := subSetID * subsetSize
	return backends[start : start+subsetSize]
}

func hashDemo(str string) int64 {
	clientHash := int64(murmur.Murmur3([]byte(str)))
	if clientHash < 0 {
		clientHash = int64(math.Abs(float64(clientHash)))
	}
	return clientHash
}
