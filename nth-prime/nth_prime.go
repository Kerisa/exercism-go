package prime


// Nth doc here
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	var primeList = []int{2}
	for cur := 3; len(primeList) < n; cur++ {
		isPrime := true
		for _,num := range primeList {
			if cur % num == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeList = append(primeList, cur)
		}
	}

	return primeList[n-1], true
}
