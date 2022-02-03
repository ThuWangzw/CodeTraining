func findMinFibonacciNumbers(k int) int {
	fibs := []int{1, 1}
	for fibs[len(fibs)-1] < k {
		fibs = append(fibs, fibs[len(fibs)-1]+fibs[len(fibs)-2])
	}
	ans := 0
	for k != 0 {
		top := fibs[len(fibs)-1]
		fibs = fibs[:len(fibs)-1]
		if top <= k {
			k -= top
			ans++
		}
	}
	return ans
}