package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := InitializeRouter()

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s...\n", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
