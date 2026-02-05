package project

import "github.com/google/uuid"

type ProjectMember struct {
	ProjectID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role      string    `gorm:"not null"`
}
