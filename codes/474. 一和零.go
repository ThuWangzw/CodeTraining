// 背包问题+状态空间压缩
// 1. 注意边界条件
// 2. 状态空间压缩时，注意遍历顺序
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMaxForm(strs []string, m int, n int) int {
	c := len(strs)
	res := make([][]int, m+1)
	counts := make([][]int, c)
	for i := 0; i < c; i++ {
		counts[i] = make([]int, 2)
		for _, cc := range strs[i] {
			if cc == '0' {
				counts[i][0] += 1
			} else {
				counts[i][1] += 1
			}
		}
	}

	for j := 0; j <= m; j++ {
		res[j] = make([]int, n+1)
	}

	for i := 0; i < c; i++ {
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if i == 0 && j >= counts[i][0] && k >= counts[i][1] {
					res[j][k] = 1
				} else if i == 0 {
					res[j][k] = 0
				} else if j >= counts[i][0] && k >= counts[i][1] {
					res[j][k] = max(res[j-counts[i][0]][k-counts[i][1]]+1, res[j][k])
				}
			}
		}
	}
	return res[m][n]
}