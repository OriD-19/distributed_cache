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

	reader := bufio.NewReader(os.Stdin)

	for {
		// wait for user input
		fmt.Print("%> ")
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))	

		if err != nil {
			fmt.Println("What the hell is this error", err)
			continue
		}

		// this line is causing a weird error
		// either the error is inside the server, or there is something wrong with this
		fmt.Println("After sending the message...")
		response, err := bufio.NewReader(conn).ReadString('\n')

		if len(message) == 0 && err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Error receiving response: ", err)
				continue
			}
			fmt.Println("error reading the response:", err)
		}

		fmt.Print(response)
	}
}
