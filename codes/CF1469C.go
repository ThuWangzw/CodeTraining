package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CF1469C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		last := []int{0, 0}
		flag := true
		for i := 0; i < n; i++ {
			gd := 0
			fmt.Fscan(in, &gd)
			if i == 0 {
				last[0], last[1] = gd, gd
			} else if i < n-1 {
				last[0] = max(last[0]-k+1, gd)
				last[1] = min(last[1]+k-1, gd+k-1)
				if last[0] > last[1] {
					flag = false
				}
			} else {
				last[0] -= k - 1
				last[1] += k - 1
				if gd < last[0] || gd > last[1] {
					flag = false
				}
			}
		}
		if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1469C(in, os.Stdout)
}
