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

	reader := bufio.NewReader(conn)

	for {
		reader.Reset(conn)
		text, err := reader.ReadString('\n')

		if len(text) == 0 && err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Error reading input:", err)
				break
			}
			continue
		}

		// split the input by whitespaces
		removeNewline := strings.Trim(text, "\n")
		words := strings.Split(removeNewline, " ")

		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]

		fmt.Println(command, args, "what?")

		// generate the command via the factory
		cmd, err := commandLine.GetCommandToExecute(command, cacheInstance, args...)

		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		msg, err := cmd.Execute()

		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		conn.Write([]byte(msg))

		if msg == "BYE" {
			return
		}
	}
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

		fmt.Println("Established a new connection")
		go handleCacheConnection(conn)
	}
}
