package hashmap

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 分析，get put时间复杂度是O(1)，迅速想到hashMap
// 最近最久未使用，用什么表示: 双向链表
type LRUCache struct {
	// 存储数据
	cache map[int]*DLinkNode

	cap  int
	size int
	// head, tail dummpy node
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func NewDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*DLinkNode)

	l := LRUCache{
		cap:   capacity,
		cache: cache,
		head:  NewDLinkNode(0, 0),
		tail:  NewDLinkNode(0, 0),
	}
	// 初始化Cache 后，确定Cache的head 和tail 互相指着
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get get value by key, return -1 if not exist key
func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	// 存在这个key，val
	node := l.cache[key]

	l.moveToHead(node) // 最近使用过，移动到链表头部
	return node.value
}

// Put 添加新的key value pair， 并把使用过的node 加到链表头部
func (l *LRUCache) Put(key, val int) {
	if _, ok := l.cache[key]; !ok {
		// 如果不存在直接加到头部, 同时需要考虑是否超出容量
		node := NewDLinkNode(key, val)
		// 添加到cache 中，差点忘记
		l.cache[key] = node
		l.addToHead(node)
		l.size++

		if l.size > l.cap {
			// 容量已经满了, 先删除尾部节点，再添加到头部
			removed := l.removeTail()
			// cache 中删除该node
			delete(l.cache, removed.key)
			l.size--
		}

	} else {
		// BUG 如果存在，先删除旧的node，再把新的node 添加到头部
		// FIX 直接更新node 的key 对应的value
		node := l.cache[key]
		node.value = val
		// 移到头部
		l.moveToHead(node)
	}
}

// removeNode 删除该节点
func (l *LRUCache) removeNode(node *DLinkNode) {
	// node.prev.next = node.next
	// node.next.prev = node.prev
	// or 分开
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// removeTail 删除尾节点
func (l *LRUCache) removeTail() *DLinkNode {
	tail := l.tail.prev

	// 删除该节点
	l.removeNode(tail)
	return tail
}

// moveToHead 移到头部 == 先删除 再 添加到头部
func (l *LRUCache) moveToHead(node *DLinkNode) {

	// 先删除该节点
	l.removeNode(node)
	// 再把该节点添加到头部
	l.addToHead(node)
}

// addToHead 添加到头部
func (l *LRUCache) addToHead(node *DLinkNode) {
	oldHead := l.head.next

	// 添加新的头节点
	l.head.next = node
	node.prev = l.head

	// 把oldHead 放到node 之后
	node.next = oldHead
	oldHead.prev = node
}
