package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	port := "8080"
	hostUrl := "127.0.1.1:" + port

	listener, err := net.Listen("tcp", hostUrl)

	if err != nil {
		fmt.Println("Error creating socket :", err)
		return
	}
	fmt.Println("server is created on port :", hostUrl)
	// close server after main exits
	defer listener.Close()
	// infinite for loop,

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error establishing connection :", err)
			continue
		}
		// fmt.Println("server is listening on port :", hostUrl)

		handleServer(conn)

	}

}

func handleServer(conn net.Conn) {
	defer conn.Close()

	str, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error Reading the message from client")
		return
	}
	fmt.Println(str)

}
