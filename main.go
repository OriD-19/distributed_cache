package main

import (
	"bufio"
	"fmt"
	"os"

	consistenthashing "github.com/OriD-19/distributed_cache/consistent_hashing"
)

func main() {
	hashRing := consistenthashing.NewSampleHashRing()

	// simple client/server implementation (not TCP yet)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a key to query: ")
		text, _ := reader.ReadString('\n')

		nodeVal := hashRing.GetNode(text)

		fmt.Println("==== Searching for values in node", nodeVal.ID)

		value, err := nodeVal.CacheClient.Get(text)

		if err != nil {
			fmt.Println("Value not found in the cache. Caching (enter a value)...")
			v, _ := reader.ReadString('\n')	
			nodeVal.CacheClient.Put(text, v)
			continue
		}

		fmt.Printf("Value found inside the cache with key %s: %s\n", text, value)
	}
}

