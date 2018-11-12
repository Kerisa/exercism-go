package accumulate

// Accumulate convert input via converter
func Accumulate(input []string, converter func(string)string) []string {
    result := make([]string, len(input))
    for i,str := range(input) {
        result[i] = converter(str)
    }
    return result
}