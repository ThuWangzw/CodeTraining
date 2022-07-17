func maximumSum(nums []int) int {
	maxMap := make(map[int][]int)
	for _, num := range nums {
		key := findBitSum(num)
		if _, ok := maxMap[key]; !ok {
			maxMap[key] = make([]int, 2)
		}
		if maxMap[key][0] <= num {
			maxMap[key][0], maxMap[key][1] = num, maxMap[key][0]
		} else if maxMap[key][1] < num {
			maxMap[key][1] = num
		}
	}
	ans := -1
	for _, pair := range maxMap {
		if pair[1] > 0 {
			ans = max(ans, pair[0]+pair[1])
		}
	}
	return ans
}

func findBitSum(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num % 10
		num /= 10
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}