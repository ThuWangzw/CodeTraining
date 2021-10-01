// 使用洗牌算法打乱数组
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{nums: nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	shuffles := make([]int, n)
	copy(shuffles, this.nums)
	for i := 0; i < n-1; i++ {
		randIdx := rand.Intn(n-i) + i
		shuffles[i], shuffles[randIdx] = shuffles[randIdx], shuffles[i]
	}
	return shuffles
}