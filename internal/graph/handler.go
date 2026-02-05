package graph

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecommendationDetail struct {
	ID             uuid.UUID `json:"id"`
	UserA          uuid.UUID `json:"user_a"`
	UserB          uuid.UUID `json:"user_b"`
	Similarity     float64   `json:"similarity"`
	CreatedAt      string    `json:"created_at"`
	SharedPatterns []string  `json:"shared_patterns"`
	TotalPatterns  int       `json:"total_patterns"`
	UserAPatterns  []string  `json:"user_a_patterns"`
	UserBPatterns  []string  `json:"user_b_patterns"`
}

func GetRecommendations(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uuid.UUID)

		var edges []UserSimilarityEdge
		db.Where("user_a = ? OR user_b = ?", userID, userID).
			Order("similarity DESC").
			Limit(5).
			Find(&edges)

		// Enrich with pattern details
		recommendations := []RecommendationDetail{}
		for _, edge := range edges {
			// Get patterns for both users
			userAPatterns := getUserPatterns(db, edge.UserA)
			userBPatterns := getUserPatterns(db, edge.UserB)

			// Find shared patterns
			sharedPatterns := findSharedPatterns(userAPatterns, userBPatterns)

			recommendations = append(recommendations, RecommendationDetail{
				ID:             edge.ID,
				UserA:          edge.UserA,
				UserB:          edge.UserB,
				Similarity:     edge.Similarity,
				CreatedAt:      edge.CreatedAt.Format("2006-01-02"),
				SharedPatterns: sharedPatterns,
				TotalPatterns:  len(sharedPatterns),
				UserAPatterns:  userAPatterns,
				UserBPatterns:  userBPatterns,
			})
		}

		c.JSON(http.StatusOK, recommendations)
	}
}

func getUserPatterns(db *gorm.DB, userID uuid.UUID) []string {
	rows, err := db.Raw(`
		SELECT DISTINCT ap.name
		FROM code_submissions cs
		JOIN submission_patterns sp ON cs.id = sp.submission_id
		JOIN algorithm_patterns ap ON sp.pattern_id = ap.id
		WHERE cs.user_id = ?
	`, userID).Rows()

	if err != nil {
		return []string{}
	}
	defer rows.Close()

	patterns := []string{}
	for rows.Next() {
		var pattern string
		rows.Scan(&pattern)
		patterns = append(patterns, pattern)
	}

	return patterns
}

func findSharedPatterns(a, b []string) []string {
	patternMap := map[string]bool{}
	for _, p := range a {
		patternMap[p] = true
	}

	shared := []string{}
	for _, p := range b {
		if patternMap[p] {
			shared = append(shared, p)
		}
	}

	return shared
}
