// LFU: Least Frequently Used，缓存满的时候，删除缓存里使用次数最少的元素，然后放入新元素，如果使用频率一样，删除缓存最久的元素
package redis

/*
基于 如果一个数据在最近一段时间内使用次数很少，那么在将来一段时间内被使用的可能性也很小
思路 利用一个数组存储数据项(cacheMap),另一个数组存放访问频次(frequentMap) 当数据项被命中时，访问频次自增，在淘汰的时候淘汰访问频次最少的数据。
	这样一来的话，在插入数据和访问数据的时候都能达到O(1)的时间复杂度，在淘汰数据的时候，通过选择算法得到应该淘汰的数据项在数组中的索引，
	并将该索引位置的内容替换为新来的数据内容即可，这样的话，淘汰数据的操作时间复杂度为O(n)。
优化
	1.引入随机算法 frequent没必要每次访问加1
	2.频率衰减 长时间未访问的降低频率
	3.初始频率提高(当前为1) 防止key刚加入就被移除
*/
// LFUCache结构：包含capacity容量, size当前容量, minFrequent当前最少访问频次, cacheMap存数据, frequentMap频次哈希表[key为访问频率value为链表]
// minFrequent为当前最少访问频次：
// 1. 插入一个新节点时，之前肯定没访问过，minFrequent = 1
// 2. set和get时，如果将key从当前频次移除后双向链表节点个数为0，且恰好是最小访问链表, minFrequent++
type LFUCache struct {
	capacity    int
	len         int
	minFrequent int
	cacheMap    map[int]*Node
	frequentMap map[int]*List
}

// 双向链表：包含head头指针, tail尾指针, len长度
type List struct {
	head *Node
	tail *Node
	len  int
}

// 节点：包含key, value, frequent访问次数, pre前驱指针, next后继指针
type Node struct {
	key      int
	value    int
	frequent int
	pre      *Node
	next     *Node
}

// addNode 双向链表辅助函数,添加一个节点到头节点后
func (l *List) addNode(node *Node) {
	head := l.head
	node.next = head.next
	node.next.pre = node
	node.pre = head
	head.next = node
}

// removeNode 双向链表辅助函数,删除一个节点
func (l *List) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// NewLFUCache 对外提供构造方法
func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity:    capacity,
		len:         0,
		minFrequent: 0,
		cacheMap:    make(map[int]*Node),
		frequentMap: make(map[int]*List),
	}
}

// fRemove 将节点从对应的频次双向链表中删除
func (l *LFUCache) fRemove(node *Node) {
	l.frequentMap[node.frequent].removeNode(node)
	l.frequentMap[node.frequent].len--
}

// fAdd 将节点添加进对应的频次双向链表，没有则创建
func (l *LFUCache) fAdd(node *Node) {
	if listNode, exist := l.frequentMap[node.frequent]; exist {
		listNode.addNode(node)
		listNode.len++
	} else {
		listNode = &List{&Node{}, &Node{}, 0}
		listNode.head.next = listNode.tail
		listNode.tail.pre = listNode.head

		listNode.addNode(node)
		listNode.len++
		l.frequentMap[node.frequent] = listNode
	}
}

// evictNode 移除一个key 同时从cacheMap和frequentMap移除
func (l *LFUCache) evictNode() {
	listNode := l.frequentMap[l.minFrequent]
	delete(l.cacheMap, listNode.tail.pre.key)
	listNode.removeNode(listNode.tail.pre)
	listNode.len--
}

// triggerVisit 触发访问(Get/Set时调用)
/*
获取一个key和修改一个key都会增加对应key的访问频次，可以独立为一个方法，完成如下任务：
1. 将对应node从频次列表中移出
2. 维护minFrequent
3. 该节点访问频次++，移动进下一个访问频次链表
*/
func (l *LFUCache) triggerVisit(node *Node) {
	l.fRemove(node)
	if node.frequent == l.minFrequent && l.frequentMap[node.frequent].len == 0 {
		l.minFrequent++
	}
	node.frequent++
	l.fAdd(node)
}

// Get 获取数据
/*
存在则增加访问频率并返回
不存在返回-1
*/
func (l *LFUCache) Get(key int) int {
	if node, exist := l.cacheMap[key]; exist {
		l.triggerVisit(node)
		return node.value
	}
	return -1
}

// Set 添加数据
/*
key存在则更新值和访问频率
不存在
	map容量到达上限则移除访问频率最低key后添加
	否则直接添加且将最低访问频率设置为1(即当前key)
*/
func (l *LFUCache) Set(key int, value int) {
	if l.capacity == 0 {
		return
	}
	if node, exist := l.cacheMap[key]; exist {
		l.triggerVisit(node)
		l.cacheMap[key].value = value
	} else {
		newNode := &Node{key, value, 1, nil, nil}
		if l.len < l.capacity {
			l.fAdd(newNode)
			l.len++
			l.minFrequent = 1
		} else {
			l.evictNode()
			l.fAdd(newNode)
			l.minFrequent = 1
		}
		l.cacheMap[key] = newNode
	}
}
