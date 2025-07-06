package main

import (
	"os"

	"github.com/OriD-19/distributed_cache/client"
	"github.com/OriD-19/distributed_cache/server"
)

func main() {
	// create a new single node TCP server
	switch os.Args[1] {
	case "server":
		server.CacheServer(8080)
	case "client":
		client.CacheClient()
	}
}

