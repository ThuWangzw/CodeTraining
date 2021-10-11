// 主要是先把问题转化为更小的数量级：拆成两段n，然后可以遍历或者双指针求解
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func factorial(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

func findPossibleSums(nums []int, k int) []int {
	n := len(nums)
	if k == 0 || n == 0 {
		return []int{0}
	}
	ans := make([]int, 0, factorial(n)/(factorial(n-k)*factorial(k)))
	if n > k {
		next := findPossibleSums(nums[1:], k)
		ans = append(ans, next...)
	}
	next := findPossibleSums(nums[1:], k-1)
	for _, pa := range next {
		ans = append(ans, pa+nums[0])
	}
	return ans
}

func minimumDifference(nums []int) int {
	n := len(nums) / 2
	a := nums[:n]
	b := nums[n:]
	sum := 0
	for _, num := range nums {
		sum += num
	}
	halfsum := sum / 2
	maxhalf := math.MinInt32
	for k := 0; k <= n; k++ {
		pa := findPossibleSums(a, k)
		sort.Slice(pa, func(i, j int) bool {
			return pa[i] < pa[j]
		})
		pb := findPossibleSums(b, n-k)
		sort.Slice(pb, func(i, j int) bool {
			return pb[i] < pb[j]
		})
		j := 0
		for i := 0; i < len(pa); i++ {
			for j < len(pb) && (pa[i]+pb[j] <= halfsum) {
				j++
			}
			if j == len(pb) {
				j--
			}
			for j >= 0 && pa[i]+pb[j] > halfsum {
				j--
			}
			if j < 0 {
				break
			}
			maxhalf = max(maxhalf, pa[i]+pb[j])
		}
	}
	return int(math.Abs(float64(sum - maxhalf*2)))
}