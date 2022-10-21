/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	q, p := l.Next, l
	for q != nil {
		newq := q.Next
		q.Next = p
		p, q = q, newq
	}
	l.Next = nil
	return p
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := reverseList(l1), reverseList(l2)
	var h, t *ListNode
	last := 0
	for p != nil && q != nil {
		node := &ListNode{
			Val: (p.Val + q.Val + last) % 10,
		}
		last = (p.Val + q.Val + last) / 10
		if h == nil {
			h, t = node, node
		} else {
			t.Next = node
			t = node
		}
		p, q = p.Next, q.Next
	}

	if q != nil {
		p = q
	}
	for p != nil {
		node := &ListNode{
			Val: (p.Val + last) % 10,
		}
		last = (p.Val + last) / 10
		t.Next = node
		t = node
		p = p.Next
	}
	if last != 0 {
		t.Next = &ListNode{
			Val: last,
		}
		t = t.Next
	}
	t.Next = nil
	return reverseList(h)
}