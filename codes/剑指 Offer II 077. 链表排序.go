/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head.Next
	slow := head
	var lastslow *ListNode
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
		lastslow = slow
		slow = slow.Next
	}
	if lastslow != nil {
		lastslow.Next = nil
	}
	left := sortList(head)
	right := sortList(slow)
	var mergedhead *ListNode
	var p *ListNode
	for left != nil && right != nil {
		next := left
		if left.Val < right.Val {
			next = left
			left = left.Next
		} else {
			next = right
			right = right.Next
		}
		if p == nil {
			mergedhead, p = next, next
			p.Next = nil
		} else {
			p.Next = next
			p = next
			p.Next = nil
		}
	}
	for left != nil {
		p.Next = left
		p = left
		left = left.Next
		p.Next = nil
	}
	for right != nil {
		p.Next = right
		p = right
		right = right.Next
		p.Next = nil
	}
	return mergedhead
}