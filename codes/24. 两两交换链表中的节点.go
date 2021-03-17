// 链表问题用递归，可以使得代码更简洁
func swapPairs(head *ListNode) *ListNode {
	if (head == nil) || (head.Next==nil) {
		return head
	}
	p := head
	q := head.Next
	p.Next = swapPairs(q.Next)
	q.Next = p
	return q
}

//用迭代的方法会更快
func swapPairs(head *ListNode) *ListNode {
	if (head == nil) || (head.Next==nil) {
		return head
	}
	phead := &ListNode{Val:-1, Next: head}
	last := phead
	p := head
	for p!=nil {
		q := p.Next
		if q == nil {
			break
		}
		p.Next = q.Next
		q.Next = p
		last.Next = q
		last = p
		p = p.Next
	}
	return phead.Next
}