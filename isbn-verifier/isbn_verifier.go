package isbn

import (
	"strings"
	"unicode"
	"strconv"
)

// IsValidISBN doc here
func IsValidISBN(isbn string) bool {
	isbn = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == 'X' || r == 'x' {
			return r
		}
		return -1
	}, isbn)
	
	const LENGTH int = 10

	if len(isbn) != LENGTH {
		return false
	}

	var sum int
	for i,ch := range isbn {
		val,ok := strconv.Atoi(string(ch))
		if ok == nil {
			sum += (LENGTH-i) * val
		} else if i == LENGTH - 1 {
			sum += 10
		} else {
			return false
		}
	}
	return (sum % 11 == 0)
}