package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CF954C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n int
	nums := make([]int, 0)
	margin := 0
	maxv := 0
	for fmt.Fscan(in, &n); n > 0; n-- {
		var v int
		fmt.Fscan(in, &v)
		maxv = max(maxv, v)
		nums = append(nums, v)
		if len(nums) >= 2 {
			delta := abs(nums[len(nums)-1] - nums[len(nums)-2])
			if delta == 0 {
				fmt.Fprintln(out, "NO")
				return
			}
			if delta > 1 && margin == 0 {
				margin = delta
			} else if delta > 1 && margin != 0 && margin != delta {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	if margin == 0 {
		fmt.Fprintln(out, "YES")
		fmt.Fprintf(out, "1 %d\n", 1000000000)
		return
	}
	for i := 1; i < len(nums); i++ {
		if abs(nums[i]-nums[i-1]) == 1 {
			if (nums[i]-1)/margin != (nums[i-1]-1)/margin {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}
	fmt.Fprintln(out, "YES")
	fmt.Fprintf(out, "%d %d\n", 1000000000, margin)
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF954C(in, os.Stdout)
}
