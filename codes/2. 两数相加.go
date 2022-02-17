func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	last := 0
	for l1 != nil && l2 != nil {
		node := &ListNode{
			Val:  (last + l1.Val + l2.Val) % 10,
			Next: nil,
		}
		last = (last + l1.Val + l2.Val) / 10
		if head == nil {
			head, tail = node, node
		} else {
			tail.Next = node
			tail = node
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		l1 = l2
	}
	for l1 != nil {
		node := &ListNode{
			Val:  (last + l1.Val) % 10,
			Next: nil,
		}
		last = (last + l1.Val) / 10
		tail.Next = node
		tail = node
		l1 = l1.Next
	}
	if last > 0 {
		node := &ListNode{
			Val:  last,
			Next: nil,
		}
		tail.Next = node
		tail = node
	}
	return head
}
