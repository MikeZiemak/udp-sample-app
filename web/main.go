package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"time"
)

type Message struct {
	Content string
}

func udpClient(message string) (string, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", "server:8080")
	if err != nil {
		return "", fmt.Errorf("error resolving server address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return "", fmt.Errorf("error connecting to UDP server: %v", err)
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("error sending message: %v", err)
	}

	buf := make([]byte, 1024)

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	response := string(buf[:n])
	return response, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		message := r.FormValue("message")
		response, err := udpClient(message)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, Message{Content: response})
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	} else {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting web server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

