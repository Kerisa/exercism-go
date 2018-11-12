package lsproduct

import (
	"errors"
	"regexp"
)

func productNums(b []byte) int64 {
	var result int64 = 1
	for _,ch := range b {
		result *= int64((ch - '0'))
	}
	return result
}

// LargestSeriesProduct doc here
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	} else if span < 0 || span > len(digits) {
		return 0, errors.New("invalid span")
	}

	re := regexp.MustCompile("^[0-9]+$")
	if !re.MatchString(digits) {
		return 0, errors.New("invalid digits")
	}

	var result, p int64
	b := []byte(digits)
	for i := 0; i < len(b) - span + 1; i++ {
		p = productNums(b[i:i+span])
		if p > result {
			result = p
		}
	}
	return result, nil
}
