package dto

type UserRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
