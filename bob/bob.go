package bob

import (
    "strings"
    "unicode"
)

// ContainLetter should have a comment documenting it.
func ContainLetter(remark string) bool {
    for _,s := range(remark) {
        if unicode.IsLetter(rune(s)) {
            return true
        }
    }
    return false
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
    trimed := strings.TrimSpace(remark)
    if len(trimed) == 0 {
        return "Fine. Be that way!"
    }

    question := trimed[len(trimed) - 1] == '?'
    yell     := (strings.ToUpper(trimed) == trimed) && ContainLetter(trimed)

    if yell && question {
        return "Calm down, I know what I'm doing!"
    } else if yell {
        return "Whoa, chill out!"
    } else if question {
        return "Sure."
    } else {
        return "Whatever."
    }
}
