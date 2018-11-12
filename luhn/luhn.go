package luhn

import (
	"unicode"
	"strings"
)

func removeWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// Valid doc here
func Valid(input string) bool {
	input = removeWhiteSpace(input)
	if len(input) <= 1 {
		return false
	}
	
	// from right to left
	var sum int
	for i,ch := range input {
		if !unicode.IsDigit(ch) {
			return false
		}

		intVal := int(ch - '0')
		if (len(input) - i) % 2 == 0 {
			intVal = (2 * intVal)
			if intVal > 9 {
				intVal -= 9
			}
		}
		sum += intVal
	}

	return (sum % 10 == 0)
}
