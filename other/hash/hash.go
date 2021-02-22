package hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

/*
一致性Hash算法Go语言简单实现
*/

type Hash func(data []byte) uint32
type Map struct {
	hash     Hash           // 计算 hash 的函数
	replicas int            // 副本数(一个物理节点对应多少个虚拟节点)
	keys     []int          // 有序的列表，从大到小排序的，这个很重要
	hashMap  map[int]string // 可以理解成用来记录虚拟节点和物理节点元数据关系的
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		// 默认可以用 crc32 来计算hash值
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 增加物理节点（可以使用IP或hostname做key）
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 将 str(num)+key 作为虚拟节点的Key
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			// map来存储 虚拟节点到物理节点 映射关系
			m.hashMap[hash] = key
		}
	}
	// 升序排列 确保后面Get()方法能找到正确的节点
	sort.Ints(m.keys)
}

// Get 根据用户输入Key确定在哪个物理节点
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	// 根据用户输入key值，计算出一个hash值
	hash := int(m.hash([]byte(key)))
	// 使用Hash值进行比较,确定属于哪个虚拟节点
	// 这里使用的是 m.keys[i] >= hash 来判断属于哪个节点，如果keys是无序的就可能会出现误差
	// 不过按理来说升序应该是最简单的
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	if idx == len(m.keys) {
		idx = 0
	}
	// 选择到对应物理节点
	return m.hashMap[m.keys[idx]]
}
func (m *Map) IsEmpty() bool {
	return len(m.hashMap) == 0
}
