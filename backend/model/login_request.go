package model

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,excludesall= "`
	Password string `json:"password" validate:"required,alphanum,excludesall= "`
}
