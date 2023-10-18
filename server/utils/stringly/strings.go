package stringly

import (
	"strings"
	"unicode"
)

func ToSnakeCase(input string) string {
	var result strings.Builder

	// Iterate through each character in the input string
	for i, char := range input {
		if i > 0 && unicode.IsUpper(char) {
			// If the character is an uppercase letter and not the first character, add an underscore
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(char))
	}

	return result.String()
}
