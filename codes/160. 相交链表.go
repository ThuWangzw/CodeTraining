// 双指针法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB
	for p != q {
		if p == nil {
			p = headB
		} else if q == nil {
			q = headA
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return p
}