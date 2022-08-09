package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func CF467C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	nums := make([]int64, n)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
		prefix[i+1] = prefix[i] + nums[i]
	}
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, k+1)
	}
	for i := m; i <= n; i++ {
		for j := 1; j <= k; j++ {
			var ans int64 = 0
			if i >= 1 {
				ans = dp[i-1][j]
			}
			if i >= m && j >= 1 {
				ans = max(ans, dp[i-m][j-1]+prefix[i]-prefix[i-m])
			} else if i >= m && j == 0 {
				ans = max(ans, prefix[i]-prefix[i-m])
			} else if i == m-1 && j == 0 {
				ans = max(ans, prefix[i])
			}
			max(dp[i-1][j], dp[i-m][j-1]+prefix[i]-prefix[i-m])
			dp[i][j] = ans
		}
	}
	fmt.Fprintln(out, dp[n][k])
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF467C(in, os.Stdout)
}
