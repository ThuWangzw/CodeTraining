package main

import (
	"bufio"
	"fmt"
	"os"
)

func CF280B(nums []int) int {
	st := make([]int, 0)
	ans := 0
	for _, num := range nums {
		for len(st) > 0 && st[len(st)-1] < num {
			ans = max(ans, num^st[len(st)-1])
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = max(ans, num^st[len(st)-1])
		}
		st = append(st, num)
	}
	return ans
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
	var n int
	fmt.Fscan(in, &n)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &as[i])
	}
	fmt.Fprintln(out, CF280B(as))
}
