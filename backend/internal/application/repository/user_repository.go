package repository

import (
	"canary-stream/backend/internal/application/repository/query"
	"canary-stream/backend/internal/domain"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func (repository *userRepository) InsertRegister(ctx context.Context, username string, hash string) (bool, error) {
	query := fmt.Sprintf(
		query.UserCreateNew,
		username,
		hash,
		domain.UserAdmin,
		domain.UserListener,
		domain.StatusActive,
		domain.StatusPending,
	)

	if _, err := repository.db.Exec(ctx, query); err != nil {
		slog.Error("Error inserting new user",
			"event", "db.exec_query",
			"repository", "user.insert_register",
			"username", username,
			"hash", hash,
			"user_admin", domain.UserAdmin,
			"user_listener", domain.UserListener,
			"status_active", domain.StatusActive,
			"status_pending", domain.StatusPending,
			"error", err,
		)

		return false, err
	}

	slog.Info("New user inserted successful",
		"event", "repository.insert_register",
		"repository", "user.insert_register",
	)

	return true, nil
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
