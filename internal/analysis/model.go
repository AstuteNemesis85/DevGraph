package analysis

import (
	"time"

	"github.com/google/uuid"
)

type AlgorithmPattern struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"unique;not null"`
}

type CodeAnalysis struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	SubmissionID    uuid.UUID `gorm:"type:uuid;not null;index"`
	TimeComplexity  string
	SpaceComplexity string
	Issues          string
	CreatedAt       time.Time
}
