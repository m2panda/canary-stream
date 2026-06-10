package usecase

import (
	"canary-stream/backend/internal/domain"
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repository domain.UserRepository
}

func (usecase *userUseCase) CreateRegister(ctx context.Context, username string, password string) (bool, error) {
	var bytePassword []byte = []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing user password: %v", err)
		return false, fmt.Errorf("Error hashing password")
	}

	return usecase.repository.InsertRegister(ctx, username, string(hash))
}

func NewUserUseCase(repository domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repository: repository}
}
