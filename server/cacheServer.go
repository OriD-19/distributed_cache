package server

import (
	"fmt"
	"net"
)

func handleCacheConnection(conn net.Conn) {
	defer conn.Close()
}

func CacheServer(port int) {

	// listen to TPC connections
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)	
	}

	defer ln.Close()

	// accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleCacheConnection(conn)
	}
}
