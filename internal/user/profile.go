package user

import "github.com/google/uuid"

type UserAlgorithmProfile struct {
	UserID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	PatternID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	ConfidenceScore int       `gorm:"not null"`
}
