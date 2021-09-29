// 方法一：艾式筛,从素数的角度出发，遇到一个素数，就把它的倍数标记成合数
func countPrimes(n int) int {
	count := 0
	notprime := make([]bool, n)
	if n <= 2 {
		return 0
	}
	notprime[2] = false
	for i := 2; i < n; i++ {
		if !notprime[i] {
			count++
			for k := i; k*i < n; k++ {
				notprime[k*i] = true
			}
		}
	}
	return count
}

// 方法二：线性筛, 对于每个数，它总被它最小的质因数筛去
func countPrimes(n int) int {
	primes := make([]int, 0)
	notprime := make([]bool, n)
	if n <= 2 {
		return 0
	}
	for i := 2; i < n; i++ {
		if !notprime[i] {
			primes = append(primes, i)
		}
		for k := 0; k < len(primes) && primes[k]*i < n; k++ {
			notprime[primes[k]*i] = true
			if i%primes[k] == 0 {
				break
			}
		}
	}
	return len(primes)
}