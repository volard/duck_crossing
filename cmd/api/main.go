package main

import (
	"log"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
