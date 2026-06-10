package repository

import (
	"canary-stream/backend/internal/application/repository/query"
	"canary-stream/backend/internal/domain"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func (repository *userRepository) InsertRegister(ctx context.Context, username string, hash string) (bool, error) {
	var query string = fmt.Sprintf(
		query.UserCreateNew,
		username,
		hash,
		domain.UserAdmin,
		domain.UserListener,
		domain.StatusActive,
		domain.StatusPending,
	)

	log.Print(query)

	_, err := repository.db.Exec(ctx, query)

	if err != nil {
		log.Printf("Failed to insert new user: %v", err)
		return false, fmt.Errorf("Error inserting new user")
	}

	return true, nil
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
