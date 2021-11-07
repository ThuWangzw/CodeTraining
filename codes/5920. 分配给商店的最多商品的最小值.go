func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Piece struct {
	after int
	cut   int
	num   int
}

type PriorityQueue struct {
	nums []Piece
}

func Constructor(nums []int) *PriorityQueue {
	i := len(nums) - 1
	for i > 0 {
		if nums[i] > nums[(i-1)/2] {
			nums[i], nums[(i-1)/2] = nums[(i-1)/2], nums[i]
		}
		i--
	}
	pieces := make([]Piece, 0)
	pq := &PriorityQueue{nums: pieces}

	for i := 0; i < len(nums); i++ {
		piece := Piece{after: nums[i], cut: 1, num: nums[i]}
		pq.Insert(piece)
	}
	return pq
}

func (pq *PriorityQueue) up() {
	i := len(pq.nums) - 1
	for i > 0 {
		if pq.nums[i].after > pq.nums[(i-1)/2].after {
			pq.nums[i], pq.nums[(i-1)/2] = pq.nums[(i-1)/2], pq.nums[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
}

func (pq *PriorityQueue) down() {
	n := len(pq.nums)
	i := 0
	for i < n {
		maxidx := i
		if i*2+1 < n && pq.nums[i*2+1].after > pq.nums[i].after {
			maxidx = i*2 + 1
		}
		if i*2+2 < n && pq.nums[i*2+2].after > pq.nums[maxidx].after {
			maxidx = i*2 + 2
		}
		if maxidx == i {
			break
		}
		pq.nums[i], pq.nums[maxidx] = pq.nums[maxidx], pq.nums[i]
		i = maxidx
	}

}
func (pq *PriorityQueue) Insert(piece Piece) {
	pq.nums = append(pq.nums, piece)
	pq.up()
}

func (pq *PriorityQueue) Pop() Piece {
	n := len(pq.nums)
	ans := pq.nums[0]
	pq.nums[0], pq.nums[n-1] = pq.nums[n-1], pq.nums[0]
	pq.nums = pq.nums[:n-1]
	pq.down()
	return ans
}

func (pq *PriorityQueue) Top() Piece {
	return pq.nums[0]
}

func minimizedMaximum(n int, quantities []int) int {
	m := len(quantities)
	pq := Constructor(quantities)
	for m < n {
		top := pq.Pop()
		top.cut++
		top.after = top.num / top.cut
		if top.num%top.cut != 0 {
			top.after++
		}
		pq.Insert(top)
		m++
	}
	return pq.Top().after
}