package reverse

import "bytes"

// String doc here
func String(input string) string {
    var buffer bytes.Buffer
    R := []rune(input)
    for i := len(R) - 1; i >= 0; i-- {
        buffer.WriteString(string(R[i]))
    }
    return buffer.String()
}
