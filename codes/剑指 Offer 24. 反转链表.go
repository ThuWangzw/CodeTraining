/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	q := head.Next
	for q != nil {
		q.Next, q, p = p, q.Next, q
	}
	head.Next = nil
	return p
}