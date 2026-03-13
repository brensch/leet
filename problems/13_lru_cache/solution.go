package lru_cache

// LRUCache stores key/value pairs with least-recently-used eviction.
type LRUCache struct{}

// Constructor initializes an LRUCache with the provided capacity.
func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

// Get returns the value for key, or -1 if it does not exist.
func (c *LRUCache) Get(key int) int {
	return -1
}

// Put inserts or updates a key/value pair in the cache.
func (c *LRUCache) Put(key int, value int) {}
