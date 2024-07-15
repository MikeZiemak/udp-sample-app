package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func main() {
	host := GetEnv("HOST_NAME", "localhost")
	port, err := strconv.Atoi(GetEnv("BACKEND_PORT", "8080"))
	if err != nil {
		log.Fatalf("Error fetching port from ENV: %v", err)
	}

	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(host),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()
	fmt.Printf("UDP server listening on port %v\n", port)

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

