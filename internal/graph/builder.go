package graph

import "gorm.io/gorm"

func BuildUserProfiles(db *gorm.DB) ([]UserPatternProfile, error) {
	rows, err := db.Raw(`
		SELECT cs.user_id, ap.name
		FROM code_submissions cs
		JOIN submission_patterns sp ON cs.id = sp.submission_id
		JOIN algorithm_patterns ap ON sp.pattern_id = ap.id
	`).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := map[string]map[string]int{}

	for rows.Next() {
		var userID string
		var pattern string

		rows.Scan(&userID, &pattern)

		if _, ok := profiles[userID]; !ok {
			profiles[userID] = map[string]int{}
		}
		profiles[userID][pattern]++
	}

	result := []UserPatternProfile{}
	for uid, patterns := range profiles {
		result = append(result, UserPatternProfile{
			UserID:   uid,
			Patterns: patterns,
		})
	}

	return result, nil
}
