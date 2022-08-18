package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CF1032C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	as := make([]int, n)
	bs := make([]int, n)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, 5)
		fmt.Fscan(in, &as[i])
		if i == 0 {
			for j := 0; j < 5; j++ {
				dp[i][j] = true
			}
		} else {
			if as[i] > as[i-1] {
				for k := 0; k < 5; k++ {
					if dp[i-1][k] {
						for t := k + 1; t < 5; t++ {
							dp[i][t] = true
						}
					}
				}
			} else if as[i] < as[i-1] {
				for k := 0; k < 5; k++ {
					if dp[i-1][k] {
						for t := 0; t < k; t++ {
							dp[i][t] = true
						}
					}
				}
			} else {
				for k := 0; k < 5; k++ {
					if dp[i-1][k] {
						for t := 0; t < 5; t++ {
							if t != k {
								dp[i][t] = true
							}
						}
					}
				}
			}
			fail := true
			for k := 0; k < 5; k++ {
				if dp[i][k] {
					fail = false
					break
				}
			}
			if fail {
				fmt.Fprintln(out, "-1")
				return
			}
		}
	}
	for k := 0; k < 5; k++ {
		if dp[n-1][k] {
			bs[n-1] = k
			break
		}
	}
	for i := n - 2; i >= 0; i-- {
		if as[i] < as[i+1] {
			for k := 0; k < bs[i+1]; k++ {
				if dp[i][k] {
					bs[i] = k
					break
				}
			}
		} else if as[i] > as[i+1] {
			for k := bs[i+1] + 1; k < 5; k++ {
				if dp[i][k] {
					bs[i] = k
					break
				}
			}
		} else {
			for k := 0; k < 5; k++ {
				if dp[i][k] && k != bs[i+1] {
					bs[i] = k
					break
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", bs[i]+1)
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1032C(in, os.Stdout)
}
