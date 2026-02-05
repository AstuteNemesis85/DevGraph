package graph

import (
	"time"

	"github.com/google/uuid"
)

type UserSimilarityEdge struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	UserA      uuid.UUID `gorm:"index"`
	UserB      uuid.UUID `gorm:"index"`
	Similarity float64
	CreatedAt  time.Time
}
