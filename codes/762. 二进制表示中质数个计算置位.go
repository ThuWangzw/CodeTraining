var primes map[int]struct{} = map[int]struct{}{2: {}, 3: {}, 5: {}, 7: {}, 11: {}, 13: {}, 17: {}, 19: {}, 23: {}, 29: {}, 31: {}}

func isPrimeSet(num int) bool {
	cnt := 0
	for num > 0 {
		num = num & (num - 1)
		cnt++
	}
	_, ok := primes[cnt]
	return ok
}

func countPrimeSetBits(left int, right int) int {
	cnt := 0
	for i := left; i <= right; i++ {
		if isPrimeSet(i) {
			cnt++
		}
	}
	return cnt
}