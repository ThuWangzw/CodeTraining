/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	last := head
	p := head.Next
	for p != nil {
		if p.Val == val {
			last.Next = p.Next
			break
		}
		last, p = p, p.Next
	}
	return head
}