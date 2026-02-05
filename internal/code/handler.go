package code

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"devgraph/internal/analysis"

)

func SubmitCode(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SubmitCodeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userIDRaw, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		userID := userIDRaw.(uuid.UUID)

		submission := CodeSubmission{
			ID:         uuid.New(),
			UserID:     userID,
			Language:   req.Language,
			SourceCode: req.SourceCode,
			CreatedAt:  time.Now(),
		}

		if err := db.Create(&submission).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store submission"})
			return
		}
		analysis.JobQueue <- submission.ID


		c.JSON(http.StatusCreated, gin.H{
			"submission_id": submission.ID,
			"message":       "code submitted successfully",
		})
	}
}
