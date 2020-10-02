package Leetcode

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

//双端队列节点
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

//双端队列节点构造函数
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

//初始化一个空的LRUCache
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lc *LRUCache) Get(key int) int {
	if _, ok := lc.cache[key]; !ok {
		return -1
	}
	node := lc.cache[key]
	lc.moveToHead(node)
	return node.value
}

func (lc *LRUCache) Put(key int, value int) {
	if _, ok := lc.cache[key]; !ok {
		node := initDLinkedNode(key, value) //根据新的值创建双端队列节点
		lc.cache[key] = node
		lc.addToHead(node)
		lc.size++
		if lc.size > lc.capacity {
			removed := lc.removeTail()
			delete(lc.cache, removed.key)
			lc.size--
		}
	} else {
		node := lc.cache[key]
		node.value = value
		lc.moveToHead(node)
	}
}

func (lc *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
}

func (lc *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lc *LRUCache) moveToHead(node *DLinkedNode) {
	lc.removeNode(node)
	lc.addToHead(node)
}

func (lc *LRUCache) removeTail() *DLinkedNode {
	node := lc.tail.prev
	lc.removeNode(node)
	return node
}
