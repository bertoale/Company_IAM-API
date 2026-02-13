package auth

import "company_iam/internal/user"

type LoginRequest struct {
	Identifier string `json:"identifier"` // email atau username
	Password   string `json:"password"`
}

type LoginResponse struct {
	ID           uint   `json:"id"`
	Identifier   string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SessionResponse struct {
	User         *user.UserResponse `json:"user"`
	Roles        []string           `json:"roles"`
	Permissions  []string           `json:"permissions"`
	Applications []string           `json:"applications"`
}
