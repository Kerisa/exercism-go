package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram  map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA        []byte

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
    count := 0
    for _,b := range(d) {
        if nucleotide == b {
            count++
        }
    }
	return count, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = make(Histogram)

	for _,n := range(d) {
        if n != 'A' && n != 'C' && n != 'G' && n != 'T' {
            return h, errors.New("unknwon nucleotide")
        }
	}

	for _,n := range([]byte{'A', 'C', 'G', 'T'}) {
        if count,err := d.Count(n); err == nil {
           h[n] = count
        } else {
            return h,err
        }
    }
	return h, nil
}
