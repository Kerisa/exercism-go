
package acronym

import "strings"
import "regexp"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

    result := ""
    re := regexp.MustCompile("[^a-zA-Z]+")
    for _,subStr := range(re.Split(s, -1)) {
        result += strings.ToUpper(string(subStr[0]))
    }

	return result
}
