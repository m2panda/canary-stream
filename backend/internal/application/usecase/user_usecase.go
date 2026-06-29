package usecase

import (
	"canary-stream/backend/internal/domain"
	"context"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repository domain.UserRepository
}

func (usecase *userUseCase) CreateRegister(ctx context.Context, username string, password string) (bool, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		slog.Error("Error hashing user password",
			"event", "bcrypt.password_hashing",
			"usecase", "user.create_register",
			"error", err,
		)

		return false, err
	}

	slog.Info("Password hashed successful",
		"event", "bcrypt.password_hashing",
		"usecase", "user.create_register",
	)

	return usecase.repository.InsertRegister(ctx, username, string(hash))
}

func NewUserUseCase(repository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repository: repository}
}
