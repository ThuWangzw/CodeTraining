type PriorityQueue []*ListNode

func (pq *PriorityQueue) down(i int, n int) {
	least := i
	left := i*2 + 1
	if left < n && (*pq)[i].Val > (*pq)[left].Val {
		least = left
	}
	right := i*2 + 2
	if right < n && (*pq)[least].Val > (*pq)[right].Val {
		least = right
	}
	if least != i {
		(*pq)[i], (*pq)[least] = (*pq)[least], (*pq)[i]
		pq.down(least, n)
	}
}

func (pq *PriorityQueue) up() {
	n := len(*pq)
	i := n - 1
	for i > 0 {
		p := (i - 1) / 2
		if (*pq)[i].Val < (*pq)[p].Val {
			(*pq)[i], (*pq)[p] = (*pq)[p], (*pq)[i]
			i = p
		} else {
			return
		}
	}
}

func NewPriorityQueue(lists []*ListNode) *PriorityQueue {
	pq := PriorityQueue(lists)
	n := len(pq)
	for i := n - 1; i > 0; i-- {
		p := (i - 1) / 2
		if pq[p].Val > pq[i].Val {
			pq[p], pq[i] = pq[i], pq[p]
			pq.down(i, n)
		}
	}
	return &pq
}

func (pq *PriorityQueue) Top() *ListNode {
	return (*pq)[0]
}

func (pq *PriorityQueue) Pop() *ListNode {
	top := pq.Top()
	n := len(*pq)
	(*pq)[0], (*pq)[n-1] = (*pq)[n-1], (*pq)[0]
	*pq = (*pq)[:n-1]
	pq.down(0, n-1)
	return top
}

func (pq *PriorityQueue) Push(node *ListNode) {
	(*pq) = append((*pq), node)
	pq.up()
}

func mergeKLists(lists []*ListNode) *ListNode {
	// filter out empty list
	n := len(lists)
	filteredList := make([]*ListNode, 0, n)
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			filteredList = append(filteredList, lists[i])
		}
	}
	pq := NewPriorityQueue(filteredList)
	var begin *ListNode
	var p *ListNode
	for len(*pq) > 0 {
		top := pq.Pop()
		if begin == nil {
			begin = top
			p = top
			top = top.Next
			p.Next = nil
		} else {
			p.Next = top
			top = top.Next
			p = p.Next
			p.Next = nil
		}
		if top != nil {
			pq.Push(top)
		}
	}
	return begin
}