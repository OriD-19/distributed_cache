package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// create a simple TCP client to interact with the node instance

func CacheClient() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Println("Connected to Cache Instance in port 8080")

	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("%> ")
	for reader.Scan() {
		// wait for user input

		message := reader.Text()

		_, err := conn.Write([]byte(message + "\n"))	

		if err != nil {
			fmt.Println("What the hell is this error", err)
			continue
		}

		// this line is causing a weird error
		// either the error is inside the server, or there is something wrong with this
		
		serverReader := bufio.NewScanner(conn)
		serverReader.Scan()
		response := serverReader.Text()

		fmt.Println(response)
		
		if message == "EXIT" {
			break
		}

		fmt.Print("%> ")
	}
}
