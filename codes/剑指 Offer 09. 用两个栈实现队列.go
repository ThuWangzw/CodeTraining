type Stack []int

func (stack *Stack) Push(i int) {
	(*stack) = append((*stack), i)
}

func (stack *Stack) Pop() int {
	n := len(*stack)
	if n == 0 {
		return -1
	}
	top := (*stack)[n-1]
	(*stack) = (*stack)[:n-1]
	return top
}

func (stack *Stack) Len() int {
	return len(*stack)
}

type CQueue struct {
	Positive Stack
	Negative Stack
}

func Constructor() CQueue {
	return CQueue{
		Positive: Stack(make([]int, 0)),
		Negative: Stack(make([]int, 0)),
	}
}

func (this *CQueue) AppendTail(value int) {
	for this.Negative.Len() > 0 {
		this.Positive.Push(this.Negative.Pop())
	}
	this.Positive.Push(value)
}

func (this *CQueue) DeleteHead() int {
	for this.Positive.Len() > 0 {
		this.Negative.Push(this.Positive.Pop())
	}
	return this.Negative.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */