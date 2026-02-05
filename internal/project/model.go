package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
}
