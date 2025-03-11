package network

import (
	"bufio"
	"fmt"
	"net"
)

func StartTCPServer(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer listener.Close()

	fmt.Println("TCP Server started on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Received:", scanner.Text())
	}
}
