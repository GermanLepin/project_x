package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	UserType  string    `json:"user_type"`
}

type LoginResponse struct {
	SessionID             uuid.UUID `json:"session_id"`
	IsBlocked             bool      `json:"is_bloked"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	UserID                uuid.UUID `json:"user_id"`
}

type RefreshTokenResponse struct {
	SessionID             uuid.UUID `json:"session_id"`
	IsBlocked             bool      `json:"is_bloked"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	UserID                uuid.UUID `json:"user_id"`
}

type FetchUserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	UserType  string    `json:"user_type"`
}

type DeleteUserResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
