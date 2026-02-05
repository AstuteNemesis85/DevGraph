package analysis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnalysisResponse struct {
	ID              uuid.UUID `json:"id"`
	SubmissionID    uuid.UUID `json:"submission_id"`
	TimeComplexity  string    `json:"time_complexity"`
	SpaceComplexity string    `json:"space_complexity"`
	Issues          string    `json:"issues"`
	Patterns        []string  `json:"patterns"`
	CreatedAt       string    `json:"created_at"`
}

func GetAnalysis(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		submissionID := c.Param("id")

		var analysis CodeAnalysis
		if err := db.Where("submission_id = ?", submissionID).First(&analysis).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "analysis not found"})
			return
		}

		// Get patterns
		var patterns []AlgorithmPattern
		db.Raw(`
			SELECT ap.* FROM algorithm_patterns ap
			JOIN submission_patterns sp ON sp.pattern_id = ap.id
			WHERE sp.submission_id = ?
		`, submissionID).Scan(&patterns)

		patternNames := make([]string, len(patterns))
		for i, p := range patterns {
			patternNames[i] = p.Name
		}

		response := AnalysisResponse{
			ID:              analysis.ID,
			SubmissionID:    analysis.SubmissionID,
			TimeComplexity:  analysis.TimeComplexity,
			SpaceComplexity: analysis.SpaceComplexity,
			Issues:          analysis.Issues,
			Patterns:        patternNames,
			CreatedAt:       analysis.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		c.JSON(http.StatusOK, response)
	}
}

// Get all submissions for a user
func GetUserSubmissions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uuid.UUID)

		var submissions []struct {
			ID         uuid.UUID `json:"id"`
			Language   string    `json:"language"`
			SourceCode string    `json:"source_code"`
			CreatedAt  string    `json:"created_at"`
		}

		db.Raw(`
			SELECT id, language, source_code, created_at
			FROM code_submissions
			WHERE user_id = ?
			ORDER BY created_at DESC
		`, userID).Scan(&submissions)

		c.JSON(http.StatusOK, submissions)
	}
}
