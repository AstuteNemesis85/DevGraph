package auth

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;index"`
	RefreshTokenHash string    `gorm:"not null"`
	ExpiresAt       time.Time `gorm:"not null"`
	CreatedAt       time.Time
}
