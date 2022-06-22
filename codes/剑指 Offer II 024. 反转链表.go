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
	last := head
	cur := head.Next
	for cur != nil {
		cur, cur.Next, last = cur.Next, last, cur
	}
	head.Next = nil
	return last
}