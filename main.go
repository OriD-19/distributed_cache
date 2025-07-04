package main

import (
	"fmt"

	consistenthashing "github.com/OriD-19/distributed_cache/consistent_hashing"
)

func main() {
	hashRing := consistenthashing.NewSampleHashRing()

	keySample := "PruebaCarlitos"

	nodeVal := hashRing.GetNode(keySample)

	fmt.Println(nodeVal.ID)
}

