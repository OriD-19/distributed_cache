package consistenthashing

import (
	"fmt"
	"hash/maphash"
	"slices"

	"github.com/OriD-19/distributed_cache/lruCache"
)

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
	Ring map[uint64]*CacheNode
	Nodes    []uint64 // sorted list with all the hashes of nodes
	Hasher maphash.Hash
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

	hashRing.Ring = make(map[uint64]*CacheNode)
	hashRing.Nodes = []uint64{}
	var h maphash.Hash
	hashRing.Hasher = h

	return &hashRing
}

func (hr *HashRing) getValueHash(val string) uint64 {
	hr.Hasher.WriteString(val)
	value := hr.Hasher.Sum64()

	hr.Hasher.Reset()
	return value
}

// function for finding the position of the next node inside the ring.
// see the functionality of the consistent hashing algorithm with 
// clockwise policy
func (hr *HashRing) binarySearchNode(hash uint64) uint64 {
	
	l := 0
	r := len(hr.Nodes) - 1
	m := 0
	finishOnRight := 0

	for l <= r {

		m = (l + r) / 2
		
		if hr.Nodes[m] == hash {
			return hr.Nodes[m]
		} else if hr.Nodes[m] > hash {
			r = m - 1
			finishOnRight = 0
		} else if hr.Nodes[m] < hash {
			l = m + 1
			finishOnRight = 1
		}
	}

	// return the hash associated with the next Node that suits the current hash
	return hr.Nodes[uint64((m + finishOnRight) % len(hr.Nodes))]
}

func (hr *HashRing) InsertNode(node *CacheNode) uint64 {
	nodeHash := hr.getValueHash(node.ID)
	
	hr.Ring[nodeHash] = node
	hr.Nodes = append(hr.Nodes, nodeHash)
	// sort it for maintaning order inside the Ring
	slices.Sort(hr.Nodes)

	fmt.Println("New node added:", nodeHash)

	return nodeHash
}

func (hr *HashRing) RemoveNode(nodeRef *CacheNode) uint64 {

	nodeHash := hr.getValueHash(nodeRef.ID)

	delete(hr.Ring, nodeHash)

	// delete the node inside the list of hashes (Go 1.21+)
	hr.Nodes = slices.DeleteFunc(hr.Nodes, func(val uint64) bool {
		return val == nodeHash
	})
	
	fmt.Println("Node removed:", nodeHash)

	return nodeHash
}

func (hr *HashRing) GetNode(key string) *CacheNode {
	
	keyHash := hr.getValueHash(key)

	// perform binary search through all the node hashes registered
	nodeHash := hr.binarySearchNode(keyHash)

	return hr.Ring[nodeHash]
}
