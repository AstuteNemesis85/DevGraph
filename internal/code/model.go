package code

import (
	"time"

	"github.com/google/uuid"
)

type CodeSubmission struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Language   string    `gorm:"not null"`
	SourceCode string    `gorm:"type:text;not null"`
	CreatedAt  time.Time
}
