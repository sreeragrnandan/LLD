package lrucache

import (
	"fmt"
	"sync"
)

// Node ...
type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

// LRUCache ...
type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	mu       sync.RWMutex
}

// NewLRUCache ...
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	cache := &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
	}
	cache.head = &Node[K, V]{}
	cache.tail = &Node[K, V]{}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (c *LRUCache[K, V]) Get(key K) (value V, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if node, found := c.cache[key]; found {
		c.moveToHead(node)
		return node.value, true
	}

	return
}

func (c *LRUCache[K, V]) GetCache() {
	fmt.Printf("%+v", c.cache)
	return
}

func (c *LRUCache[K, V]) Put(key K, value V) {

	if node, found := c.cache[key]; found {
		node.value = value
		c.moveToHead(node)
		return
	}
	node := &Node[K, V]{
		key:   key,
		value: value,
	}
	c.cache[key] = node
	// add to head
	c.addToHead(node)

	if len(c.cache) > c.capacity {
		// remove from tail
		tailNode := c.removeTail(node)
		delete(c.cache, tailNode.key)
	}
	return
}

func (c *LRUCache[K, V]) addToHead(node *Node[K, V]) {

	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache[K, V]) remove(node *Node[K, V]) {

	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache[K, V]) moveToHead(node *Node[K, V]) {
	c.remove(node)
	c.addToHead(node)
}

func (c *LRUCache[K, V]) removeTail(node *Node[K, V]) *Node[K, V] {
	node = c.tail.prev
	c.remove(node)
	return node
}

func (c *LRUCache[K, V]) Clear(node *Node[K, V]) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[K]*Node[K, V])
	c.head.next = c.tail
	c.tail.prev = c.head
}
