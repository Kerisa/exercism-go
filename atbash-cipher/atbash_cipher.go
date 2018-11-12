package atbash

import (
	"strings"
)

// Atbash makes atbash cipher
func Atbash(input string) string {
	var result string
	wrote := 0
	for _,c := range strings.ToLower(input) {
		if c >= 'a' && c <= 'z' {
			c = 'z' - c + 'a'
		} else if c >= '0' && c <= '9'{
		} else {
			continue
		}
		if wrote > 0 && wrote % 5 == 0 {
			result += " "
		}
		result += string(c)
		wrote++
	}
	return result
}