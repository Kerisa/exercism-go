package etl

import "strings"

// Transform doc here
func Transform(input map[int][]string) map[string]int {
	var result = map[string]int{}
	for score,group := range input {
		for _,ch:= range group {
			result[strings.ToLower(ch)] = score
		}
	}
	return result
}
