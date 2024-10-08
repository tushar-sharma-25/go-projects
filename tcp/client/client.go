package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	port := "8080"
	hostUrl := "127.0.1.1:" + port

	conn, err := net.Dial("tcp", hostUrl)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error connecting to the server :", err)
		return
	}

	fmt.Println("Client is created :")

	handleClient(conn)
	fmt.Scanln()

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	rw := bufio.NewReadWriter(reader, writer)

	rw.WriteString("Hello from client")

	rw.Flush()

	message, err := rw.ReadString('\n')

	if err != nil {
		fmt.Println("Error recieving message from the server :", err)
		return
	}

	fmt.Println("message from server", message)

}
