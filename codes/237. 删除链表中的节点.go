func deleteNode(node *ListNode) {
	for {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			return
		}
		node = node.Next
	}
}