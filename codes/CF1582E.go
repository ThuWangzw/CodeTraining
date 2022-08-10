package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MAXK = 500

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CF1582E(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		nums := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[n-i-1])
		}
		presum := make([]int64, n+1)
		for i := 1; i <= n; i++ {
			presum[i] = presum[i-1] + nums[i-1]
		}

		dp := make([][]int64, n+1)
		dp[0] = make([]int64, MAXK)
		maxk := 0
		for i := 1; i <= n; i++ {
			dp[i] = make([]int64, MAXK)
			for j := 1; j <= MAXK-1; j++ {
				dp[i][j] = dp[i-1][j]
				if i >= j && (j == 1 || (dp[i-j][j-1] > 0 && presum[i]-presum[i-j] < dp[i-j][j-1])) && (presum[i]-presum[i-j] > dp[i-1][j]) {
					dp[i][j] = presum[i] - presum[i-j]
				}
				if dp[i][j] > 0 {
					maxk = max(maxk, j)
				}
			}
		}
		fmt.Fprintln(out, maxk)
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1582E(in, os.Stdout)
}
