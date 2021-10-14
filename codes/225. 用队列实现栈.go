type MyStack struct {
	positive []int
	negative []int
}

func Constructor() MyStack {
	return MyStack{positive: make([]int, 0), negative: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.positive = append(this.positive, x)
}

func (this *MyStack) Pop() int {
	stackLen := len(this.positive)
	for i := 0; i < stackLen-1; i++ {
		this.negative = append(this.negative, this.positive[0])
		this.positive = this.positive[1:]
	}
	top := this.positive[0]
	this.positive = this.negative
	this.negative = make([]int, 0)
	return top
}

func (this *MyStack) Top() int {
	stackLen := len(this.positive)
	for i := 0; i < stackLen-1; i++ {
		this.negative = append(this.negative, this.positive[0])
		this.positive = this.positive[1:]
	}
	top := this.positive[0]
	this.negative = append(this.negative, this.positive[0])
	this.positive = this.negative
	this.negative = make([]int, 0)
	return top
}

func (this *MyStack) Empty() bool {
	return (len(this.positive) == 0) && (len(this.negative) == 0)
}