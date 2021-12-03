func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func largestSumAfterKNegations(nums []int, k int) int {
	numMap := make([]int, 201)
	sum := 0
	for _, num := range nums {
		numMap[num+100] += 1
		sum += num
	}
	for i := 0; i < 100; i++ {
		if k == 0 {
			break
		}
		if numMap[i] == 0 {
			continue
		}
		if numMap[i] <= k {

		}
		delcnt := min(k, numMap[i])
		sum += (100 - i) * 2 * delcnt
		k -= delcnt
	}
	if k%2 == 1 {
		for i := 0; i <= 100; i++ {
			if numMap[100+i] != 0 || numMap[100-i] != 0 {
				sum -= i * 2
				break
			}
		}
	}
	return sum
}