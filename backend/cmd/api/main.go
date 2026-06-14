package main

import (
	"fmt"
	"log/slog"
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
	core.InitLoggerService()

	if err := core.RegisterCustomValidators(); err != nil {
		slog.Error("Error register custom validators",
			"event", "validator.custom_validators",
			"status", 500,
			"error", err,
		)

		return
	}

	server := http.NewServeMux()

	apiPort := os.Getenv("API_PORT")
	muxPort := fmt.Sprintf(":%s", apiPort)

	if err := framework.RouterSetup(server); err != nil {
		slog.Error("Error setup server router for api rest",
			"event", "router.setup",
			"status", 500,
			"error", err,
		)

		return
	}

	if err := http.ListenAndServe(muxPort, server); err != nil {
		slog.Error("Error launching mux server",
			"event", "server.listen_start",
			"port", apiPort,
			"status", 500,
			"error", err,
		)
		return
	}
}
