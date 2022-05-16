func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	n := len(postorder)
	mid := -1
	for i := 0; i < n-1; i++ {
		if mid == -1 && postorder[i] > postorder[n-1] {
			mid = i
		}
		if postorder[i] < postorder[n-1] && mid > -1 {
			return false
		}
	}
	if mid == -1 {
		return verifyPostorder(postorder[:n-1])
	}
	return verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:n-1])
}