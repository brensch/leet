package lru_cache

// node stores one cache entry inside the recency list.
type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

// LRUCache stores key/value pairs with least-recently-used eviction.
type LRUCache struct {
	capacity int
	items    map[int]*node
	head     *node
	tail     *node
}

// Constructor initializes an LRUCache with the provided capacity.
func Constructor(capacity int) LRUCache {
	// Sentinel nodes remove most nil edge cases when inserting/removing.
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		items:    map[int]*node{},
		head:     head,
		tail:     tail,
	}
}

// Get returns the value for key, or -1 if it does not exist.
func (c *LRUCache) Get(key int) int {
	entry, ok := c.items[key]
	if !ok {
		return -1
	}

	// Access changes recency, so move this node to the front.
	c.remove(entry)
	c.insertAfterHead(entry)

	return entry.value
}

// Put inserts or updates a key/value pair in the cache.
func (c *LRUCache) Put(key int, value int) {
	if entry, ok := c.items[key]; ok {
		// Updating an existing key should not evict anything.
		entry.value = value
		c.remove(entry)
		c.insertAfterHead(entry)
		return
	}

	if len(c.items) == c.capacity {
		// The real LRU item is always the node right before tail.
		lru := c.tail.prev
		c.remove(lru)
		delete(c.items, lru.key)
	}

	entry := &node{key: key, value: value}
	c.items[key] = entry
	c.insertAfterHead(entry)
}

func (c *LRUCache) remove(entry *node) {
	entry.prev.next = entry.next
	entry.next.prev = entry.prev
}

func (c *LRUCache) insertAfterHead(entry *node) {
	entry.prev = c.head
	entry.next = c.head.next
	c.head.next.prev = entry
	c.head.next = entry
}
