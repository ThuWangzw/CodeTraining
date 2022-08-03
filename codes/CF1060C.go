package main

import (
	"bufio"
	"fmt"
	"os"
)

func CF1060C(as []int, bs []int, x int) int {
	m := len(as)
	n := len(bs)
	presuma := make([]int, m)
	presumb := make([]int, n)
	presuma[0] = as[0]
	presumb[0] = bs[0]
	for i := 1; i < m; i++ {
		presuma[i] = presuma[i-1] + as[i]
	}
	for i := 1; i < n; i++ {
		presumb[i] = presumb[i-1] + bs[i]
	}
	mina := make([]int, m)
	minb := make([]int, n)
	for i := 1; i <= m; i++ {
		sum := presuma[i-1]
		for j := i; j < m; j++ {
			sum = min(sum, presuma[j]-presuma[j-i])
		}
		mina[i-1] = sum
	}
	for i := 1; i <= n; i++ {
		sum := presumb[i-1]
		for j := i; j < n; j++ {
			sum = min(sum, presumb[j]-presumb[j-i])
		}
		minb[i-1] = sum
	}
	maxarea := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if int64(mina[i])*int64(minb[j]) <= int64(x) {
				maxarea = max(maxarea, (i+1)*(j+1))
			}
		}
	}
	return maxarea
}

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

func main() {
	// read, _ := os.Open("input.txt")
	read := os.Stdin
	write := os.Stdout
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()
	var n, m, x int
	fmt.Fscan(in, &n, &m)
	as := make([]int, n)
	bs := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &bs[i])
	}
	fmt.Fscan(in, &x)
	fmt.Fprintln(out, CF1060C(as, bs, x))
}
