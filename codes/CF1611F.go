package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1611F(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var n int
		var s int64
		fmt.Fscan(in, &n, &s)
		nums := make([]int, n)
		var sum int64
		left := 0
		l, r := -1, -2
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[i])
			sum += int64(nums[i])
			for sum+s < 0 && left <= i {
				sum -= int64(nums[left])
				left++
			}
			if sum+s >= 0 && r-l < i-left {
				l, r = left+1, i+1
			}
		}
		if l == -1 {
			fmt.Fprintln(out, l)
		} else {
			fmt.Fprintln(out, l, r)
		}
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1611F(in, os.Stdout)
}
