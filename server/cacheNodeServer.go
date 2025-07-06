package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/OriD-19/distributed_cache/commandLine"
	"github.com/OriD-19/distributed_cache/lruCache"
)

/*
This simple TCP server will handle cache connections,
as well as executing commands that interact with the Cache instance via TCP
*/

var cacheInstance *lruCache.LRUCache

func handleCacheConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewScanner(conn)

	for reader.Scan() {
		text := reader.Text()

		// split the input by whitespaces
		words := strings.Split(text, " ")

		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]

		// generate the command via the factory
		cmd, err := commandLine.GetCommandToExecute(command, cacheInstance, args...)

		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			continue
		}

		msg, err := cmd.Execute()

		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			continue
		}

		conn.Write([]byte(msg + "\n"))

		if msg == "BYE" {
			break
		}
	}

	if reader.Err() != nil {
		fmt.Println("Error while reading input:", reader.Err())
	}

	fmt.Println("Connection closed for address", conn.RemoteAddr())
}

func CacheServer(port int) {

	// listen to TPC connections
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	defer ln.Close()

	cacheInstance = lruCache.NewLRUCache(100)
	// accept connections
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
