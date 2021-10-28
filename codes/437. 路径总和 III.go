func pathSumIter(root *TreeNode, targetSum int) (int, []int) {
	if root == nil {
		return 0, make([]int, 0)
	}
	lcnt, lsums := pathSumIter(root.Left, targetSum)
	rcnt, rsums := pathSumIter(root.Right, targetSum)
	cnt := lcnt + rcnt
	sums := append(lsums, rsums...)
	for i := 0; i < len(sums); i++ {
		sums[i] += root.Val
		if sums[i] == targetSum {
			cnt++
		}
	}
	sums = append(sums, root.Val)
	if root.Val == targetSum {
		cnt++
	}
	return cnt, sums
}

func pathSum(root *TreeNode, targetSum int) int {
	cnt, _ := pathSumIter(root, targetSum)
	return cnt
}