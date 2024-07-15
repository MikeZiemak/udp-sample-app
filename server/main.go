package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("server"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP server listening on port 8080")

	buf := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}
		message := string(buf[:n])
		fmt.Printf("Received message from %s: %s\n", remoteAddr, message)

		response := "Echo: " + message
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}
}

