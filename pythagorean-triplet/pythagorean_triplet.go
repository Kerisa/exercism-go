package pythagorean

import "math"

// Triplet doc here
type Triplet [3]int

// Range doc here
func Range(min, max int) []Triplet {
	var list []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c := math.Sqrt(float64(a * a + b * b))
			if int(c) <= max && math.Ceil(c) == c {
				list = append(list, Triplet{a, b, int(c)})
			}
		}
	}
	return list
}

// Sum doc here
func Sum(p int) []Triplet {
	var list []Triplet
	for a := 1; a < p; a++ {
		for b := a + 1; b < p - a; b++ {
			c := p - a - b
			if a * a + b * b == c * c {
				list = append(list, Triplet{a, b, c})
			}
		}
	}
	return list
}