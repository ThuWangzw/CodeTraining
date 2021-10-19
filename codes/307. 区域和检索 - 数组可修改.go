// 线段树
type NumArray struct {
	left       int
	right      int
	leftChild  *NumArray
	rightChild *NumArray
	val        int
}

func Constructor(nums []int) NumArray {
	// init nodes
	n := len(nums)
	nodes := make([]NumArray, n)
	for i := 0; i < n; i++ {
		nodes[i] = NumArray{left: i, right: i, leftChild: nil, rightChild: nil, val: nums[i]}
	}
	for len(nodes) > 1 {
		newnodes := make([]NumArray, 0)
		i := 0
		for i = 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) {
				newnodes = append(newnodes, NumArray{left: nodes[i].left, right: nodes[i].right,
					leftChild: &nodes[i], rightChild: nil, val: nodes[i].val})
			} else {
				newnodes = append(newnodes, NumArray{left: nodes[i].left, right: nodes[i+1].right,
					leftChild: &nodes[i], rightChild: &nodes[i+1], val: nodes[i].val + nodes[i+1].val})
			}
		}
		nodes = newnodes
	}
	return nodes[0]
}

func (this *NumArray) Update(index int, val int) {
	if this.leftChild == nil {
		this.val = val
		return
	} else if index <= this.leftChild.right {
		this.leftChild.Update(index, val)
	} else {
		this.rightChild.Update(index, val)
	}
	if this.rightChild == nil {
		this.val = this.leftChild.val
	} else {
		this.val = this.leftChild.val + this.rightChild.val
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if this.left == left && this.right == right {
		return this.val
	} else if right <= this.leftChild.right {
		return this.leftChild.SumRange(left, right)
	} else if left >= this.rightChild.left {
		return this.rightChild.SumRange(left, right)
	} else {
		return this.leftChild.SumRange(left, this.leftChild.right) + this.rightChild.SumRange(this.rightChild.left, right)
	}
}