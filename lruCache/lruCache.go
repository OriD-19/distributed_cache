package lruCache

import "fmt"

type LRUCache struct {
	List    *DoublyLinkedList
	HashMap map[string]*Node
	Capacity int
}

func NewLRUCache(cap int) *LRUCache {

	var cache LRUCache

	cache.List = NewDoublyLinkedList(nil)
	cache.HashMap = make(map[string]*Node)
	cache.Capacity = cap

	return &cache
}

func (cache *LRUCache) Put(key, value string) error {

	// check if we need to free space

	if cache.List.Size == cache.Capacity {
		// evict the least recently used item
		ref, err := cache.List.RemoveAtTail()
		if err != nil {
			return fmt.Errorf("could not insert in cache: %v", err)
		}
		// remove the record inside the hashmap as well
		delete(cache.HashMap, ref.Data.Key)
	}

	nodeRef := cache.List.InsertAtHead(&key, &value)	
	cache.HashMap[key] = nodeRef

	return nil
}

func (cache *LRUCache) Get(key string) (string, error) {

	nodeRef, ok := cache.HashMap[key]

	if !ok {
		return "", fmt.Errorf("cache miss: key not found")
	}

	// reinster the node as a high proiority node
	// spwan a goroutine? interesting
	cache.List.ReInsertAtHead(nodeRef)
	
	return *nodeRef.Data.Value, nil	
}
