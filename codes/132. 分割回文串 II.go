// 动态规划，注意一个缩减时间的方法是：计算回文串可以缩减到O1，因为可以预先计算
func minCut(s string) int {
	n := len(s)
	result := make([]int, n, n)
	result[0] = 0
	// find is palidrome
	ispalindrome := make([][]bool, 0, n)
	for i:=0; i<n; i++ {
		ispalindrome = append(ispalindrome, make([]bool, n, n))
	}
	for i:=n-1; i>=0; i-- {	
		for j:=0; j<n; j++ {
			if i>=j {
				ispalindrome[i][j] = true
			} else {
				ispalindrome[i][j] = ispalindrome[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	for i:=1; i<n; i++ {
		result[i] = result[i-1] + 1
		for j:=i-1; j>=0; j-- {
			if ispalindrome[j][i] {
				if j==0 {
					result[i] = 0
				} else {
					result[i] = int(math.Min(float64(result[i]), float64(result[j-1]+1)))
				}
			}
		}
	}
	return result[n-1]
}