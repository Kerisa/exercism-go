package wordcount

import (
	"strings"
	"unicode"
)

// Frequency doc here
type Frequency map[string]int

// WordCount doc here
func WordCount(input string) Frequency {
	var result = Frequency{}
	group := strings.FieldsFunc(input, func(ch rune)bool{
		return !(unicode.IsLetter(ch) || unicode.IsNumber(ch) || ch == '\'')
	})
	for _,str := range group {
		result[strings.ToLower(strings.Trim(str, "'"))]++
	}
	return result
}
