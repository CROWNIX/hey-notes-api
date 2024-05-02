package dto

import "hey-notes-api/models"

type RegisterRequest struct {
	Username        string `json:"username" validate:"required,max=32"`
	Email           string `json:"email" validate:"required,email,max=64"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  *models.User
	Token string `json:"token"`
}