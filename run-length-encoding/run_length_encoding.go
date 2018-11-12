package encode

import (
	"bytes"
	"strconv"
	"strings"
)

func encodeMultiChar(ch rune, count int) string {
	var result string
	if count > 0 {
		if count > 1 {
			result += strconv.Itoa(count)
		}
		result += string(ch)
	}
	return result
}

// RunLengthEncode doc here
func RunLengthEncode(input string) string {
	var encodeStr string
	
	var cache rune
	var count int
	for _,ch := range input {
		if ch == cache {
			count++
		} else {
			encodeStr += encodeMultiChar(cache, count)
			count = 1
			cache = ch
		}
	}
	encodeStr += encodeMultiChar(cache, count)
	return encodeStr
}

// RunLengthDecode doc here
func RunLengthDecode(input string) string {
	var decodeStr bytes.Buffer
	var cacheNumber bytes.Buffer
	
	for _,ch := range input {
		if (ch >= '0' && ch <= '9') {
			cacheNumber.WriteString(string(ch))
		} else {
			if cacheNumber.Len() == 0 {
				decodeStr.WriteString(string(ch))
			} else {
				number,_ := strconv.Atoi(cacheNumber.String())
				decodeStr.WriteString(strings.Repeat(string(ch), number))
				cacheNumber.Reset()
			}
		}
	}
	return decodeStr.String()
}
