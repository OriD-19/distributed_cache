package server

import (
	"fmt"
	"net"

	consistenthashing "github.com/OriD-19/distributed_cache/consistent_hashing"
)

var distributedCache *consistenthashing.HashRing

func handleCacheProxyConnection(conn net.Conn) {
	defer conn.Close()

}

func CacheProxyServer(port int) {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	// ideally, we would start this thing with a better setup...
	// like health checks, monitoring, node addition strategies configured, etc.
	distributedCache = consistenthashing.NewSampleHashRing()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Printf("Client %s connected\n", conn.RemoteAddr())
		go handleCacheConnection(conn)
	}
	
}
