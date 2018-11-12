package strain

// Ints type...
type Ints               []int
// Lists type...
type Lists              [][]int
// Strings type...
type Strings            []string

type intPredFunc        func(x int) bool
type intArrayPredFunc   func(x []int) bool
type strPredFunc        func(x string) bool

// Keep Ints Keep
func (L *Ints) Keep(pred intPredFunc) Ints {
	var result Ints
	for _, ii := range *L {
		if pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}

// Discard Ints Discard
func (L *Ints) Discard(pred intPredFunc) Ints {
	var result Ints
	for _, ii := range *L {
		if !pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}

// Keep Lists Keep
func (L *Lists) Keep(pred intArrayPredFunc) Lists {
	var result Lists
	for _, ii := range *L {
		if pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}

// Discard Lists Discard
func (L *Lists) Discard(pred intArrayPredFunc) Lists {
	var result Lists
	for _, ii := range *L {
		if !pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}

// Keep Strings Keep
func (L *Strings) Keep(pred strPredFunc) Strings {
	var result Strings
	for _, ii := range *L {
		if pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}

// Discard Strings Discard
func (L *Strings) Discard(pred strPredFunc) Strings {
	var result Strings
	for _, ii := range *L {
		if !pred(ii) {
			result = append(result, ii)
		}
	}
	return result
}
