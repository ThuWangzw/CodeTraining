/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	p := head
	q := head
	for p != nil && p.Next != nil {
		p = p.Next.Next
		if p == nil {
			break
		}
		q = q.Next
		if p == q {
			q = head
			for p != q {
				p = p.Next
				q = q.Next
			}
			return p
		}
	}
	return nil
}