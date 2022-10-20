func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	var q *ListNode
	for p != nil {
		if q == nil {
			q = head
		} else {
			q = q.Next
		}
		p = p.Next
	}
	if q == nil {
		return head.Next
	}
	q.Next = q.Next.Next
	return head
}