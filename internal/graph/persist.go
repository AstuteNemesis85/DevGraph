package graph

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PersistGraph(db *gorm.DB, edges []UserSimilarity) error {
	for _, e := range edges {
		userAID, err := uuid.Parse(e.UserA)
		if err != nil {
			return err
		}

		userBID, err := uuid.Parse(e.UserB)
		if err != nil {
			return err
		}

		edge := UserSimilarityEdge{
			ID:         uuid.New(),
			UserA:      userAID,
			UserB:      userBID,
			Similarity: e.Similarity,
			CreatedAt:  time.Now(),
		}

		db.FirstOrCreate(
			&edge,
			UserSimilarityEdge{
				UserA: userAID,
				UserB: userBID,
			},
		)
	}
	return nil
}
