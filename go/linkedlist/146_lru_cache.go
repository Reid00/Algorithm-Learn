package linkedlist

// LRU (最近最少使用) 缓存
// 分析，get put时间复杂度是O(1)，迅速想到hashMap
// 最近最久未使用，用什么表示: 双向链表
// 双向链表 Doubly linked list
type DuLNode struct {
	// 这个里面key 可以不存在 val 一定要存在(wrong)
	// NOTE key must be exists, cache 容量超出时，要删node， 根据node.key
	key, val   int
	prev, next *DuLNode
}

func NewDuLNode(key, val int) *DuLNode {
	return &DuLNode{
		key: key,
		val: val,
	}
}

type LRUCache struct {
	// 存储数据
	cache map[int]*DuLNode

	cap  int
	size int
	// 哨兵节点
	head, tail *DuLNode
}

func Constructor(capacity int) LRUCache {
	hmap := make(map[int]*DuLNode)
	lru := LRUCache{
		cache: hmap,

		cap:  capacity,
		size: 0,
		head: NewDuLNode(0, 0),
		tail: NewDuLNode(0, 0),
	}

	// 将head 和 tail 连接起来
	lru.head.next = lru.tail
	lru.tail.prev = lru.head

	return lru
}

// Get get value from cache
// 1. key exists, get key and move node to head
// 2. key doesn't exist, return -1
func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		// not exists
		return -1
	}

	node := l.cache[key]
	// move to head
	l.moveToHead(node)
	return node.val
}

// Put put key value to cache
// 1. key exists, update value and move to head
// 2. key doesn't exist, new key,val node and put to head
func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		// 不存在 new key node and put to head
		node := NewDuLNode(key, value)
		// add to cache
		l.cache[key] = node
		// add to head
		l.addToHead(node)
		// compare size
		l.size++
		if l.size > l.cap {
			// removeTail from DuLinkList
			tail := l.removeTail()
			// delete from cache
			delete(l.cache, tail.key)
			l.size--
		}

	} else {
		// 存在，update value
		node := l.cache[key]
		node.val = value
		l.moveToHead(node)
	}
}

// removeTail remove tail node while l.size > l.cap
func (l *LRUCache) removeTail() *DuLNode {
	// dummy tail l.tail
	tail := l.tail.prev
	// remove
	l.removeNode(tail)
	return tail
}

// removeNode remove the node
func (l *LRUCache) removeNode(node *DuLNode) {

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

// moveToHead move the node to head while put method
func (l *LRUCache) moveToHead(node *DuLNode) {
	// 1. removeNode
	// 2. add node the head
	l.removeNode(node)
	l.addToHead(node)
}

// addToHead add node to head
func (l *LRUCache) addToHead(node *DuLNode) {
	// dummy node is head

	oldHead := l.head.next

	// 放到 dummy node 之后
	l.head.next = node
	node.prev = l.head
	// 放到oldHead 之前
	node.next = oldHead
	oldHead.prev = node
}
