package hamming

import "errors"

// Distance calculate distance of two string
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return -1, errors.New("not equal length")
    }

    count := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            count++
        }
    }

    return count, nil
}