package framework

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/application/repository"
	"canary-stream/backend/internal/application/usecase"
	"canary-stream/backend/internal/framework/handlers/genre"
	"canary-stream/backend/internal/framework/handlers/status"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

var mux *http.ServeMux
var db *pgxpool.Pool
var vk valkey.Client

// Support function to route status entity handlers
func statusRouter() {
	if mux == nil || db == nil || vk == nil {
		return
	}

	repository := repository.NewStatusRepository(db, vk)
	usecase := usecase.NewStatusUseCase(repository)

	getAllHandler := status.NewGetAllHandler(usecase)

	mux.Handle("GET /status", getAllHandler)
}

func genreRouter() {
	if mux == nil || db == nil {
		return
	}

	repository := repository.NewGenreRepository(db)
	usecase := usecase.NewGenreUseCase(repository)

	getAllHandler := genre.NewGetAllHandler(usecase)

	mux.Handle("GET /genres", getAllHandler)
}

/**
 * Main function to initialize server routing; first use
 * core function to connect to db and then to valkey server;
 * set package variable values; call each support function
 * to route entity endpoints
 */
func RouterSetup(server *http.ServeMux) error {
	pool, err := core.DBConnection()

	if err != nil {
		return fmt.Errorf("Error established connection with db: %w", err)
	}

	client, err := core.CacheConnection()

	if err != nil {
		return fmt.Errorf("Error established connection with valkey: %w", err)
	}

	db = pool
	vk = client
	mux = server

	statusRouter()
	genreRouter()

	return nil
}
