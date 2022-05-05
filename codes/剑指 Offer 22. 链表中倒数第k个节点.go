/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	p := head
	q := head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p != nil {
		p, q = p.Next, q.Next
	}
	return q
}