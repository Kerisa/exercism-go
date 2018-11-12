package pangram

import "strings"

// IsPangram doc here
func IsPangram(s string) bool {
	usage := map[rune]bool{}
	for _,ch := range strings.ToLower(s) {
		usage[ch] = true
	}
	for c := 'a'; c < 'z'; c++ {
		if _,ok := usage[c]; !ok {
			return false
		}
	}
	return true
}
