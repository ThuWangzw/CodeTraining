type MinStack struct {
	nums []int
	minv int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		nums: make([]int, 0),
		minv: math.MaxInt32,
	}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x-this.minv)
	if x-this.minv < 0 {
		this.minv = x
	}
}

func (this *MinStack) Pop() {
	top := this.nums[len(this.nums)-1]
	realTop := 0
	if top >= 0 {
		realTop = top + this.minv
	} else {
		realTop = this.minv
		this.minv = realTop - top
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	top := this.nums[len(this.nums)-1]
	if top >= 0 {
		return top + this.minv
	} else {
		return this.minv
	}
}

func (this *MinStack) Min() int {
	return this.minv
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */