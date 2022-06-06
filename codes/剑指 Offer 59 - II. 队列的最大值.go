type MaxQueue struct {
	queue    []int
	maxValue int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:    make([]int, 0),
		maxValue: -1,
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.maxValue
}

func (this *MaxQueue) Push_back(value int) {
	if len(this.queue) == 0 || this.maxValue < value {
		this.maxValue = value
	}
	this.queue = append(this.queue, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	top := this.queue[0]
	if top == this.maxValue {
		this.maxValue = math.MinInt32
		for i := 1; i < len(this.queue); i++ {
			this.maxValue = max(this.maxValue, this.queue[i])
		}
	}
	this.queue = this.queue[1:]
	return top
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */