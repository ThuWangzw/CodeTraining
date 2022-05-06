/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB
	if p == nil || q == nil {
		return nil
	}
	for p != q {
		p = p.Next
		q = q.Next
		if p == q && p == nil {
			return nil
		}
		if p == nil {
			p = headB
		}
		if q == nil {
			q = headA
		}
	}
	return p
}