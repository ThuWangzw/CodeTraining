func sortedListToBSTIter(head *ListNode, end *ListNode) *TreeNode {
	if head == end {
		return nil
	}
	fast := head.Next
	slow := head
	for fast != end {
		slow = slow.Next
		fast = fast.Next
		if fast != end {
			fast = fast.Next
		}
	}
	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBSTIter(head, slow),
		Right: sortedListToBSTIter(slow.Next, end),
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBSTIter(head, nil)
}