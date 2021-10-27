// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var last *ListNode
	p := head
	q := head
	t := 0
	for q.Next != nil {
		t++
		q = q.Next
		if t >= n {
			if last == nil {
				last = head
			} else {
				last = last.Next
			}
			p = p.Next
		}
	}
	if last == nil {
		return p.Next
	} else {
		last.Next = p.Next
		return head
	}
}