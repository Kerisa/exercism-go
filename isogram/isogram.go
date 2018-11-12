package isogram

import "strings"

// IsIsogram doc here
func IsIsogram(input string) bool {
    var alphaCount = map[rune]bool{}
    for _,ch := range(strings.ToLower(input)) {
        if ch >= 'a' && ch <= 'z' {
            if _,ok := alphaCount[ch]; ok {
                return false
            }
            
            alphaCount[ch] = true
        }
    }
    return true
}