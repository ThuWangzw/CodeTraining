func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	phead := head
	qhead := head.Next
	pend := head
	qend := head.Next
	p := pend
	q := qend
	for p != nil && q != nil {
		p = q.Next
		if p == nil {
			break
		}
		q = p.Next
		pend.Next = p
		qend.Next = q
		pend = p
		qend = q
	}
	pend.Next = qhead
	return phead
}