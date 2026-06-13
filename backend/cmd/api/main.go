package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"canary-stream/backend/core"
	"canary-stream/backend/internal/framework"
)

/**
 * Main function to start rest api server;
 * use env information to expore connection
 * port; create mux server and call router
 * to initialize configuration
 */
func main() {
	if err := core.RegisterCustomValidators(); err != nil {
		log.Printf("Error on register custom validators: %v", err)
		return
	}

	apiPort := os.Getenv("API_PORT")

	server := http.NewServeMux()

	if err := framework.RouterSetup(server); err != nil {
		log.Printf("Error on router setup %v", err)
		return
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", apiPort), server); err != nil {
		log.Printf("Error launching serv %v", err)
		return
	}
}
