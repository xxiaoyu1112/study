package main

import "fmt"

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func GetLruCache(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkNode{},
		capacity: capacity,
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (lc *LRUCache) Get(key int) int {
	if _, ok := lc.cache[key]; ok {
		node := lc.cache[key]
		lc.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (lc *LRUCache) Put(key, value int) {
	if _, ok := lc.cache[key]; ok {
		node := lc.cache[key]
		node.value = value
		lc.moveToHead(node)
	} else {
		node := initDLinkNode(key, value)
		lc.cache[key] = node
		lc.addToHead(node)
		lc.size++
		if lc.size > lc.capacity {
			removed := lc.GetTailNode()
			delete(lc.cache, removed.key)
			lc.size--
		}
	}
}

func (lc *LRUCache) moveToHead(node *DLinkNode) {
	lc.removeNode(node)
	lc.addToHead(node)
}

func (lc *LRUCache) addToHead(node *DLinkNode) {
	node.next = lc.head.next
	node.prev = lc.head
	lc.head.next.prev = node
	lc.head.next = node
}

func (lc *LRUCache) GetTailNode() *DLinkNode {
	node := lc.tail.prev
	lc.removeNode(node)
	return node
}

func (lc *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func main() {
	lruCache := GetLruCache(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)
	lruCache.Get(2)
	//lruCache.Put(4, 4)
	//lruCache.Get(1)
	//lruCache.Get(3)
	//lruCache.Get(4)
	fmt.Println(lruCache)
}
