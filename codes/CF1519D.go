package main

import (
	"bufio"
	"fmt"
	"os"
)

func CF1519D(as []int, bs []int) int64 {
	n := len(as)
	var ans int64
	for i := 0; i < n; i++ {
		ans += int64(as[i]) * int64(bs[i])
	}
	base := ans
	nega := make([][]int64, n)
	for i := n - 1; i >= 0; i-- {
		nega[i] = make([]int64, n)
		for j := i; j < n; j++ {
			if i == j {
				nega[i][j] = int64(as[j]) * int64(bs[j])
			} else if i == j-1 {
				nega[i][j] = int64(as[i])*int64(bs[j]) + int64(as[j])*int64(bs[i])
			} else {
				nega[i][j] = int64(as[i])*int64(bs[j]) + int64(as[j])*int64(bs[i]) + nega[i+1][j-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		var posi int64 = 0
		for j := i; j >= 0; j-- {
			posi += int64(as[j]) * int64(bs[j])
			ans = max(ans, base-posi+nega[j][i])
		}
	}
	return ans
}

func max(a, b int64) int64 {
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
	var n int
	fmt.Fscan(in, &n)
	as := make([]int, n)
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &bs[i])
	}
	fmt.Fprintln(out, CF1519D(as, bs))
}
