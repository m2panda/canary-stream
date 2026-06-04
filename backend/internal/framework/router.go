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
)

var mux *http.ServeMux
var db *pgxpool.Pool

func statusRouter() {
	if mux == nil || db == nil {
		return
	}

	repository := repository.NewStatusRepository(db)
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

func RouterSetup(server *http.ServeMux) error {
	pool, err := core.DBConnection()

	if err != nil {
		return fmt.Errorf("Error established connection with db: %w", err)
	}

	db = pool
	mux = server

	statusRouter()
	genreRouter()

	return nil
}
