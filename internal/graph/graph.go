package graph

type UserSimilarity struct {
	UserA      string
	UserB      string
	Similarity float64
}

func BuildSimilarityGraph(profiles []UserPatternProfile, threshold float64) []UserSimilarity {
	edges := []UserSimilarity{}

	for i := 0; i < len(profiles); i++ {
		for j := i + 1; j < len(profiles); j++ {
			score := WeightedJaccard(
				profiles[i].Patterns,
				profiles[j].Patterns,
			)

			if score >= threshold {
				edges = append(edges, UserSimilarity{
					UserA:      profiles[i].UserID,
					UserB:      profiles[j].UserID,
					Similarity: score,
				})
			}
		}
	}

	return edges
}
