func findTarget(root *TreeNode, k int) bool {
	nums := make([]int, 0)
	p := root
	stack := make([]*TreeNode, 0)
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, p.Val)
		p = p.Right
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == k {
			return true
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}
	return false
}