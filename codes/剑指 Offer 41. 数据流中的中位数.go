type PriorityQueue struct {
	nums    []int
	reverse bool // true for >
}

func PriorityQueueConstructor(reverse bool) PriorityQueue {
	return PriorityQueue{
		nums:    []int{},
		reverse: reverse,
	}
}

func (pq *PriorityQueue) compare(i, j int) bool {
	return (pq.nums[i] < pq.nums[j] && !pq.reverse) || (pq.nums[i] > pq.nums[j] && pq.reverse)
}

func (pq *PriorityQueue) up() {
	node := len(pq.nums) - 1
	for node > 0 {
		parent := (node - 1) / 2
		if pq.compare(node, parent) {
			pq.nums[node], pq.nums[parent] = pq.nums[parent], pq.nums[node]
			node = parent
		} else {
			return
		}
	}
}

func (pq *PriorityQueue) down(i int) {
	next := i
	if i*2+1 < len(pq.nums) && pq.compare(i*2+1, i) {
		next = i*2 + 1
	}
	if i*2+2 < len(pq.nums) && pq.compare(i*2+2, next) {
		next = i*2 + 2
	}
	if next == i {
		return
	}
	pq.nums[i], pq.nums[next] = pq.nums[next], pq.nums[i]
	pq.down(next)
}

func (pq *PriorityQueue) Push(i int) {
	pq.nums = append(pq.nums, i)
	pq.up()
}

func (pq *PriorityQueue) Pop() int {
	n := len(pq.nums) - 1
	deleted := pq.nums[0]
	pq.nums[0], pq.nums[n] = pq.nums[n], pq.nums[0]
	pq.nums = pq.nums[:n]
	pq.down(0)
	return deleted
}

func (pq *PriorityQueue) Top() int {
	return pq.nums[0]
}

func (pq *PriorityQueue) Len() int {
	return len(pq.nums)
}

type MedianFinder struct {
	leftHeap  PriorityQueue
	rightHeap PriorityQueue
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		leftHeap:  PriorityQueueConstructor(true),
		rightHeap: PriorityQueueConstructor(false),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.leftHeap.Len() == 0 {
		this.leftHeap.Push(num)
		return
	}
	leftTop := this.leftHeap.Top()
	if num < leftTop {
		this.leftHeap.Push(num)
		if this.leftHeap.Len()-this.rightHeap.Len() > 1 {
			this.rightHeap.Push(this.leftHeap.Pop())
		}
	} else {
		this.rightHeap.Push(num)
		if this.rightHeap.Len() > this.leftHeap.Len() {
			this.leftHeap.Push(this.rightHeap.Pop())
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.rightHeap.Len() == 0 {
		return float64(this.leftHeap.Top())
	} else if this.rightHeap.Len() == this.leftHeap.Len() {
		return (float64(this.rightHeap.Top()) + float64(this.leftHeap.Top())) / 2
	} else {
		return float64(this.leftHeap.Top())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */