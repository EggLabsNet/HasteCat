package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var hasteURL *string

const maxLength = 100_000 // 40KB
const version = "0.1"

func main() {

	// Define command-line options
	listenIP := flag.String("ip", "0.0.0.0", "the IP address to listen on")
	listenPort := flag.Int("port", 99, "the port number to listen on")
	hasteURL = flag.String("hasteurl", "https://haste.egglabs.net", "the URL of the Hastebin server to use")

	flag.Parse()

	log.Println("Starting HasteCat server...")

	// Listen for incoming connections
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *listenIP, *listenPort))
	if err != nil {
		log.Fatal("Error starting the server:", err.Error())
	}
	log.Printf("Server started. Listening on %s:%d\n", *listenIP, *listenPort)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}
