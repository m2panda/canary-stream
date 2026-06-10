package domain

import "context"

// Role enum type configuration
type UserRole string

const (
	UserAdmin    UserRole = "admin"
	UserManager  UserRole = "manager"
	UserListener UserRole = "listener"
)

func (role UserRole) Valid() bool {
	switch role {
	case UserAdmin, UserManager, UserListener:
		return true
	}

	return false
}

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
	Picture   string   `json:"picture"`
	Token     string   `json:"token"`
	TokenExp  string   `json:"token_exp"`
	CreatedAt string   `json:"created_at"`
	Status    Status   `json:"status"`
}

type UserRepository interface {
	InsertRegister(ctx context.Context, username string, hash string) (bool, error)
}

type UserUseCase interface {
	CreateRegister(ctx context.Context, username string, password string) (bool, error)
}
