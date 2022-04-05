func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

type SmallHeap struct {
	nums []int
}

func (heap *SmallHeap) up() {
	node := len(heap.nums) - 1
	for node > 0 {
		parent := (node - 1) / 2
		if heap.nums[parent] > heap.nums[node] {
			heap.nums[parent], heap.nums[node] = heap.nums[node], heap.nums[parent]
			node = parent
		} else {
			return
		}
	}
}

func (heap *SmallHeap) down(node int) {
	large := node
	left := node*2 + 1
	right := node*2 + 2
	n := len(heap.nums)
	if left < n && heap.nums[left] < heap.nums[node] {
		large = left
	}
	if right < n && heap.nums[right] < heap.nums[large] {
		large = right
	}
	if large != node {
		heap.nums[node], heap.nums[large] = heap.nums[large], heap.nums[node]
		heap.down(large)
	}
}

func (heap *SmallHeap) Top() int {
	return heap.nums[0]
}

func (heap *SmallHeap) Pop() int {
	n := len(heap.nums)
	num := heap.nums[0]
	heap.nums[0], heap.nums[n-1] = heap.nums[n-1], heap.nums[0]
	heap.nums = heap.nums[:n-1]
	heap.down(0)
	return num
}

func (heap *SmallHeap) Push(num int) {
	heap.nums = append(heap.nums, num)
	heap.up()
}

func (heap *SmallHeap) Size() int {
	return len(heap.nums)
}

func sortIntervals(intervals [][]int, begin, end int) {
	if begin >= end {
		return
	}
	pivot := intervals[begin][0]
	left := begin + 1
	right := end
	for left <= right {
		if intervals[left][0] <= pivot {
			left++
		} else {
			intervals[left], intervals[right] = intervals[right], intervals[left]
			right--
		}
	}
	intervals[begin], intervals[right] = intervals[right], intervals[begin]
	sortIntervals(intervals, begin, right-1)
	sortIntervals(intervals, right+1, end)
}

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	sortIntervals(intervals, 0, n-1)
	maxRooms := 0
	heap := SmallHeap{
		nums: make([]int, 0),
	}
	for _, interval := range intervals {
		for heap.Size() > 0 && heap.Top() <= interval[0] {
			heap.Pop()
		}
		heap.Push(interval[1])
		maxRooms = max(maxRooms, heap.Size())
	}
	return maxRooms
}