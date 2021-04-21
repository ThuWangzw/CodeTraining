/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	last := dummyHead
	tail := dummyHead
	p := head
	for p != nil {
		begin := p
		i := 0
		for i = 0; i < k && p != nil; i++ {
			tmp := p.Next
			p.Next = last
			if p == begin {
				p.Next = nil
			}
			last = p
			p = tmp
		}
		if i < k && p == nil {
			p = last
			last = tail
			begin = p
			tail.Next = p
			for p != nil {
				tmp := p.Next
				p.Next = last
				last = p
				p = tmp
			}
			begin.Next = nil
			tail.Next = last
			break
		}
		tail.Next = last
		tail = begin
	}
	return dummyHead.Next
}