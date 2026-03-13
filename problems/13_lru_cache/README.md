# LRU Cache

## Concept
Hash map plus doubly linked list design.

## Prompt
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the `LRUCache` type:

- `Constructor(capacity int) LRUCache` initializes the cache with a positive size capacity.
- `Get(key int) int` returns the value of the key if the key exists, otherwise returns `-1`.
- `Put(key int, value int)` updates the value of the key if it exists. Otherwise, adds the key-value pair to the cache. If the number of keys exceeds the capacity, evict the least recently used key.

Both `Get` and `Put` should run in `O(1)` average time.

## Target
Implement `Constructor`, `Get`, and `Put` in `solution.go`.
