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

type ConsistentHashing struct {
	HashRing map[int]CacheNode
	Nodes    []int // sorted list with all the hashes of nodes
}

func NewCacheNode(cache *lruCache.LRUCache)
