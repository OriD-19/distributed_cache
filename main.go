package main

import "fmt"
import "github.com/OriD-19/distributed_cache/lruCache"

func main() {

	fmt.Println("Testing LRU Cache")

	cache := lru_cache.NewLRUCache(2)

	cache.Put("f1", "Fernando")
	cache.Put("f2", "Carlitos")
	res, _ := cache.Get("f1")

	fmt.Println(res)

	cache.Put("f3", "Gerardo")

	res, _ = cache.Get("f2")

	if res == "" {
		fmt.Println("f2 not in cache anymore")
	}

	cache.Put("f4", "Carlita")

	res, _ = cache.Get("f1")
	if res == "" {
		fmt.Println("f1 not in cache anymore")
	}
	
	res, _ = cache.Get("f3")

	if res == "" {
		fmt.Println("f3 not in cache anymore")
	} else {
		fmt.Println(res)
	}

	res, _ = cache.Get("f4")

	if res == "" {
		fmt.Println("f4 not in cache anymore")
	} else {
		fmt.Println(res)
	}
}

