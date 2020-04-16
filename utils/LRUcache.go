package utils

/*
原理
使用链表 每次get/sett时将key移动到表头
这样最后在表尾的元素就是最近未使用的
实现
一个简单的node节点 用map存放数据
*/

// 双向链表结构
type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}

// LRU结构
type LRUCache struct {
	m    map[int]*LinkNode
	cap  int
	head *LinkNode
	tail *LinkNode
}

// NewLURCache 对外提供构造方法
func NewLRUCache(capacity int) LRUCache {
	// 其中head和tail节点是固定的 后续操作也不会动
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode, capacity), capacity, head, tail}
}

// moveToHead 先删除再添加到表头
func (l *LRUCache) moveToHead(node *LinkNode) {
	l.remove(node)
	l.add(node)
}

// remove 移除当前节点 修改前后节点指针即可
func (l *LRUCache) remove(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// add 在表头后新增节点
func (l *LRUCache) add(node *LinkNode) {
	head := l.head
	node.next = head.next
	node.next.pre = node
	node.pre = head
	head.next = node
}

// Get 查询
/*如果有，将这个元素移动到首位置，返回值
如果没有，返回-1*/
func (l *LRUCache) Get(key int) int {
	if v, exist := l.m[key]; exist {
		l.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

// Set 新增
/*如果元素存在，将其移动到最前面，并更新值
如果元素不存在，插入到最前面，放入map（判断容量）*/
func (l *LRUCache) Set(key int, value int) {
	tail := l.tail
	cache := l.m
	if v, exist := cache[key]; exist {
		v.value = value
		l.moveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(l.m) == l.cap {
			delete(cache, tail.pre.key)
			l.remove(tail.pre)
		}
		l.add(v)
		cache[key] = v
	}
}
