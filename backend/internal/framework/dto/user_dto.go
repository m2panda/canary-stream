package dto

type UserRegisterBody struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
	Password string `json:"password" validate:"required,min=8,securepassword"`
}
