package summultiples


// SumMultiples doc here
func SumMultiples(limit int, divisors ...int) int {
	var numberSet = map[int]bool{}
	for _,n := range divisors {		
		for ii := n; ii < limit; ii += n {
			numberSet[ii] = true
		}
	}

	var sum int
	for n := range numberSet {
		sum += n
	}
	return sum
}
