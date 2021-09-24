package hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

/*
一致性Hash算法Go语言简单实现
Consistent Hashing Go Implement
*/

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           // 计算 hash 的函数
	replicas int            // 副本数(一个物理节点对应多少个虚拟节点)
	keys     []int          // 虚拟节点计算出的 Hash 值列表,升序排列
	hashMap  map[int]string // 可以理解成用来记录虚拟节点和物理节点元数据关系的,具体为: 虚拟节点 hash 值和物理节点的映射关系
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
	// 根据用户输入 key 值，计算出一个 hash 值
	hash := int(m.hash([]byte(key)))
	// 使用 hash 值进行比较,确定属于哪个虚拟节点
	// 这里使用的是 m.keys[i] >= hash 来判断属于哪个节点，如果 keys 是无序的就可能会出现误差,所以每次添加新节点后都需要进行排序
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
