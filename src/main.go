package main

import (
	"log"
)

// Features:
// 1. When the service finds a pod that is in a Pending state it will send a nats message
func main() {
	log.Println("Starting the application...")

	for {
		// 1. Check if there are any pods in a Pending state
		// 2. If there are, send a nats message
		// 3. Wait for a few seconds before checking again
	}

}
