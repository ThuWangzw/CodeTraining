// 前缀和+二分查找
type Solution struct {
	prev []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	n := len(w)
	prev := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prev[i] = prev[i-1] + w[i-1]
	}
	return Solution{prev: prev}
}

func (this *Solution) PickIndex() int {
	n := len(this.prev)
	idx := rand.Intn(this.prev[n-1]) + 1
	start := 0
	end := n - 1
	for start < end-1 {
		mid := (start + end) / 2
		if this.prev[mid] < idx {
			start = mid
		} else {
			end = mid
		}
	}
	return start
}