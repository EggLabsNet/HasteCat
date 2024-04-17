package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var hasteURL *string

const maxLength = 100_000 // 40KB
const version = "0.1"

func main() {

	// Set some defaults
	defaultListenIP := os.Getenv("LISTEN_IP")
	if defaultListenIP == "" {
		defaultListenIP = "0.0.0.0"
	}

	var defaultListenPort int
	var err error
	defaultListenPortStr := os.Getenv("LISTEN_PORT")
	if defaultListenPortStr == "" {
		defaultListenPort = 99
	} else {
		defaultListenPort, err = strconv.Atoi(defaultListenPortStr)
		if err != nil {
			log.Fatal("Error converting listen port to integer:", err.Error())
		}
	}

	defaultHasteURL := os.Getenv("HASTEBIN_URL")
	if defaultHasteURL == "" {
		defaultHasteURL = "https://haste.egglabs.net"
	}

	// Define command-line options
	listenIP := flag.String("ip", defaultListenIP, "the IP address to listen on")
	listenPort := flag.Int("port", defaultListenPort, "the port to listen on")
	hasteURL = flag.String("hasteurl", defaultHasteURL, "the URL of the Hastebin server to use")

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
