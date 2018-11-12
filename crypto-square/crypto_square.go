package cryptosquare

import (
	"bytes"
	"unicode"
	"strings"
	"math"
)

func normalizeString(str string) string {
	str = strings.ToLower(str)
	return strings.Map(func (r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, str)
}

// Encode doc here
func Encode(plain string) string {
	plain = normalizeString(plain)
	if len(plain) == 0 {
		return ""
	}

	row := int(math.Round(math.Sqrt(float64(len(plain)))))
	col := (len(plain) + row - 1) / row
	plain += strings.Repeat(" ", row * col - len(plain))

	var result bytes.Buffer
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			result.WriteByte(plain[r * col + c])
		}

		if c < col - 1 {
			result.WriteByte(' ')
		}
	}

	return result.String()
}