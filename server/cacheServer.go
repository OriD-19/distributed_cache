package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/OriD-19/distributed_cache/commandLine"
	"github.com/OriD-19/distributed_cache/lruCache"
)

/*
This simple TCP server will handle cache connections,
as well as executing commands that interact with the Cache instance via TCP
*/

func handleCacheConnection(conn net.Conn) {
	defer conn.Close()


	reader := bufio.NewReader(conn)
	cacheInstance := lruCache.NewLRUCache(100)// capacity of 100 items per node
	
	for {
		fmt.Print("%> ")

		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// split the input by whitespaces
		words := []string{}
		for _, word := range text {
			if word != ' ' {
				words = append(words, string(word))
			}
		}

		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := words[1:]

		// generate the command via the factory
		cmd := commandLine.GetCommandToExecute(command, cacheInstance, args...)

		if cmd == nil {
			conn.Write([]byte(fmt.Sprintf("Invalid syntax for the %s command", command)))
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
