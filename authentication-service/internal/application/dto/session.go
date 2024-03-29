package dto

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID                    uuid.UUID
	UserID                uuid.UUID
	IsBlocked             bool
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
	CreatedAt             time.Time
}
