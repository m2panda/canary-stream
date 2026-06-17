package framework

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/application/repository"
	"canary-stream/backend/internal/application/usecase"
	"canary-stream/backend/internal/framework/handlers/genre"
	"canary-stream/backend/internal/framework/handlers/status"
	"canary-stream/backend/internal/framework/handlers/user"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

var mux *http.ServeMux
var db *pgxpool.Pool
var vk valkey.Client

// ! Status
func statusRouter() {
	repository := repository.NewStatusRepository(db, vk)
	usecase := usecase.NewStatusUseCase(repository)

	getAllHandler := status.NewGetAllHandler(usecase)

	mux.Handle("GET /status", getAllHandler)
}

// ! User
func userRouter() {
	repository := repository.NewUserRepository(db)
	usecase := usecase.NewUserUseCase(repository)

	createRegisterHandler := user.NewCreateRegisterHandler(usecase)

	mux.Handle("POST /register", createRegisterHandler)
}

func genreRouter() {
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
		slog.Error("Error on db connection",
			"event", "pgdb.connection",
			"error", err,
		)

		return err
	}

	client, err := core.CacheConnection()

	if err != nil {
		slog.Error("Error on cache storage connection",
			"event", "cache.connection",
			"error", err,
		)

		return err
	}

	db = pool
	vk = client
	mux = server

	if mux == nil || db == nil || vk == nil {
		slog.Error("API connection objects no available",
			"mux", mux != nil,
			"db", db != nil,
			"vk", vk != nil,
		)

		return fmt.Errorf("API connection objects no available")
	}

	statusRouter()
	userRouter()
	genreRouter()

	slog.Info("API routing complete",
		"event", "router.setup",
	)

	return nil
}
