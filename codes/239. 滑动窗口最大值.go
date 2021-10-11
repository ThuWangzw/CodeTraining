// 手撸优先队列
type NumIndex [2]int

type PriorityQueue struct {
	nums []NumIndex
}

func (this *PriorityQueue) up() {
	i := len(this.nums) - 1
	for i > 0 && this.nums[(i-1)/2][0] < this.nums[i][0] {
		this.nums[(i-1)/2], this.nums[i] = this.nums[i], this.nums[(i-1)/2]
		i = (i - 1) / 2
	}
}

func (this *PriorityQueue) down() {
	n := len(this.nums) - 1
	i := 0
	for i <= n && i*2+1 <= n {
		idx := i*2 + 1
		if idx+1 <= n && this.nums[idx][0] < this.nums[idx+1][0] {
			idx++
		}
		if this.nums[i][0] < this.nums[idx][0] {
			this.nums[i], this.nums[idx] = this.nums[idx], this.nums[i]
			i = idx
		} else {
			break
		}

	}
}

func (this *PriorityQueue) Push(num int, index int) {
	this.nums = append(this.nums, [2]int{num, index})
	this.up()
}

func (this *PriorityQueue) Pop() {
	n := len(this.nums)
	this.nums[0], this.nums[n-1] = this.nums[n-1], this.nums[0]
	this.nums = this.nums[:n-1]
	this.down()
}

func (this *PriorityQueue) Top() NumIndex {
	return this.nums[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0, n-k+1)
	pq := PriorityQueue{nums: make([]NumIndex, 0)}
	for i := 0; i < k; i++ {
		pq.Push(nums[i], i)
	}
	ans = append(ans, pq.Top()[0])
	for i := k; i < n; i++ {
		pq.Push(nums[i], i)
		for pq.Top()[1] <= i-k {
			pq.Pop()
		}
		ans = append(ans, pq.Top()[0])
	}
	return ans
}