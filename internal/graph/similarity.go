package graph

func WeightedJaccard(a, b map[string]int) float64 {
	intersection := 0
	union := 0

	visited := map[string]bool{}

	for k, v := range a {
		visited[k] = true
		if bVal, ok := b[k]; ok {
			if v < bVal {
				intersection += v
				union += bVal
			} else {
				intersection += bVal
				union += v
			}
		} else {
			union += v
		}
	}

	for k, v := range b {
		if !visited[k] {
			union += v
		}
	}

	if union == 0 {
		return 0
	}

	return float64(intersection) / float64(union)
}
