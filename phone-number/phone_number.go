package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

// Number parse whether input is a valid phone number
func Number(input string) (output string, e error) {
	output = strings.Map(func(r rune)rune{
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, input)

	if len(output) == 11 {
		if output[0] != '1' {
			e = errors.New("11 digits does not start with a 1")
		}
		output = output[1:]
	}
	
	if len(output) != 10 {
		e = errors.New("invalid length")
	} else if output[0] < '2' {
		e = errors.New("area code starts with 0 or 1")
	} else if output[3] < '2' {
		e = errors.New("exchange code starts with 0 or 1")
	}

	return
}

// AreaCode gets area code of the number
func AreaCode(input string) (output string, e error) {
	o,e := Number(input)
	return o[:3], e
}

// Format formats input number
func Format(input string) (output string, e error) {
	o,e := Number(input)
	output = fmt.Sprintf("(%s) %s-%s", o[:3], o[3:6], o[6:])
	return
}