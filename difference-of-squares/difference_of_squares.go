package diffsquares

// SumOfSquares doc here
func SumOfSquares(sequence int) int {
	return sequence * (sequence + 1) * (2 * sequence + 1) / 6
}

// SquareOfSums doc here
func SquareOfSums(sequence int) int {
	sum := (1 + sequence) * sequence / 2
	return sum * sum
}

// Difference doc here
func Difference(sequence int) int {
	// SumOfSquares is always less than SquareOfSums
	return SquareOfSums(sequence) - SumOfSquares(sequence)
}
