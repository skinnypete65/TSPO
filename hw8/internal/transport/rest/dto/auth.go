package dto

import "ecom/internal/domain"

type SignUpInput struct {
	Username string      `json:"username" validate:"required"`
	Password string      `json:"password" validate:"required"`
	Role     domain.Role `json:"role" validate:"required"`
}

type SignInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
