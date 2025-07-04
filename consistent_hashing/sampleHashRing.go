package consistenthashing

import "github.com/OriD-19/distributed_cache/lruCache"

func NewSampleHashRing() *HashRing {

	cache1 := lruCache.NewLRUCache(10)
	cache2 := lruCache.NewLRUCache(10)
	cache3 := lruCache.NewLRUCache(10)

	cacheNode1 := NewCacheNode(cache1, "Node_1")
	cacheNode2 := NewCacheNode(cache2, "Node_2")
	cacheNode3 := NewCacheNode(cache3, "Node_3")

	hashRing := NewHashRing()

	hashRing.InsertNode(cacheNode1)
	hashRing.InsertNode(cacheNode2)
	hashRing.InsertNode(cacheNode3)

	return hashRing
}
