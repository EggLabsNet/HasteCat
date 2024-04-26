package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	var data string
	defer conn.Close()

	for {
		// Set a deadline to stop reading after a certain amount of time
		// required because netcat doesn't send any kind of EOF signal
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		// Read data from the connection
		n, err := conn.Read(buffer)

		if err != nil {

			// We reached a timeout - as long as we got some data already this was a success
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// Check length of data received
				if len(data) > 0 {
					log.Printf("Data received from client %s\n", conn.RemoteAddr().String())
					break
				} else {
					log.Printf("No data received after %d seconds from client %s\n", 2, conn.RemoteAddr().String())
					return
				}
			}

			// Some other error occurred
			log.Println("Error reading:", err.Error())
			return

		}

		// Convert the received data to string
		data += string(buffer[:n])

		// Check if we reached the maximum length
		if len(data) > maxLength {
			conn.Write([]byte("Data received exceeded maximum length\n"))
			log.Printf("Data received from client %s exceeded maximum length of %d bytes \n", conn.RemoteAddr().String(), maxLength)
			return
		}

	}

	// Handle the received text and post it to hastebin
	h, err := hastePost(data, conn.RemoteAddr().String())
	if err != nil {
		conn.Write([]byte("Error posting to Hastebin\n"))
		log.Printf("Error posting to Hastebin: %s for client %s\n", err.Error(), conn.RemoteAddr().String())

		return
	}

	response := fmt.Sprintf("%s/%s\n", *hasteURL, h.Key)
	log.Printf("Hastebin document created for client %s: %s\n", conn.RemoteAddr().String(), h.Key)

	// Send the response back to the client
	conn.Write([]byte(response))
}
