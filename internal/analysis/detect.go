package analysis

import "strings"

func detectPatterns(code string) []string {
	patterns := []string{}

	if strings.Contains(code, "for") || strings.Contains(code, "while") {
		patterns = append(patterns, "Loop")
	}
	if strings.Count(code, "for")+strings.Count(code, "while") >= 2 {
		patterns = append(patterns, "Nested Loop")
	}
	if strings.Contains(code, "map") || strings.Contains(code, "unordered_map") {
		patterns = append(patterns, "Hashing")
	}
	if strings.Contains(code, "sort(") {
		patterns = append(patterns, "Sorting")
	}

	return patterns
}
