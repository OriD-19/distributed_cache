package consistenthashing

import "github.com/OriD-19/distributed_cache/lruCache"

type CacheNodeStatus int

const (
	RUNNING CacheNodeStatus = iota
	DOWN
	MAINTENANCE
)

type CacheNode struct {
	ID          string
	Status      CacheNodeStatus
	CacheClient *lruCache.LRUCache
}

// Consistent Hashing implementation
type HashRing struct {
	Ring map[int]*CacheNode
	Nodes    []int // sorted list with all the hashes of nodes
}

func NewCacheNode(cache *lruCache.LRUCache, id string) *CacheNode {
	
	var cacheNode CacheNode

	cacheNode.Status = RUNNING // running node by default
	cacheNode.CacheClient = cache
	cacheNode.ID = id

	return &cacheNode
}

func NewHashRing() *HashRing {

	var hashRing HashRing	

	hashRing.Ring = make(map[int]*CacheNode)
	hashRing.Nodes = []int{}

	return &hashRing
}
