package anagram

import (
	"sort"
	"strings"
)

func sortStringByCharacterInLower(s string) string {
	var r []rune
	for _, runeValue := range strings.ToLower(s) {
			r = append(r, runeValue)
	}
	sort.Slice(r, func(i,j int)bool {
		return r[i] < r[j]
	})
	
	return string(r)
}

// Detect doc here
func Detect(subject string, candidates []string) []string {
	result := []string{}
	
	orignalLength := len(subject)
	lowerString	  := strings.ToLower(subject)
	orignalChars  := sortStringByCharacterInLower(subject)
	for _,str := range candidates {
		if len(str) != orignalLength || strings.ToLower(str) == lowerString {
			continue
		}
		if orignalChars == sortStringByCharacterInLower(str) {
			result = append(result, str)
		}
	}
	return result
}
