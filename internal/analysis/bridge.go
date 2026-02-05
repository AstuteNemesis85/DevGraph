package analysis

import "github.com/google/uuid"

type SubmissionPattern struct {
	SubmissionID uuid.UUID `gorm:"type:uuid;primaryKey"`
	PatternID    uuid.UUID `gorm:"type:uuid;primaryKey"`
}
