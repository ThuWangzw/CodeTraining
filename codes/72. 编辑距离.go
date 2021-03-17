// 与最长公共子序列非常像
func min(a int, b int, c int) int {
	tmp := a
	if tmp > b {
		tmp = b
	}
	if tmp > c {
		tmp = c
	}
	return tmp
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	result := make([][]int, 0, m+1)
	for i:=0; i<=m; i++ {
		result = append(result, make([]int, n+1, n+1))
	}
	//init
	for j:=1; j<=n; j++{
		result[0][j] = j
	}
	for i:=1; i<=m; i++ {
		result[i][0] = i
	}
	for i:=1; i<=m; i++ {	
		for j:=1; j<=n; j++ {
			flag := 1
			if word1[i-1]==word2[j-1] {
				flag = 0
			}
			result[i][j] = min(result[i-1][j]+1, result[i][j-1]+1, result[i-1][j-1]+flag)
		}
	}
	return result[m][n]
}