package analysis

import "strings"

func inferComplexity(code string) (string, string) {
	loops := strings.Count(code, "for") + strings.Count(code, "while")

	switch loops {
	case 0:
		return "O(1)", "O(1)"
	case 1:
		return "O(n)", "O(1)"
	default:
		return "O(n^2)", "O(1)"
	}
}
