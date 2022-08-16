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

func CF1249E(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n, c int
	fmt.Fscan(in, &n, &c)
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &b[i])
	}

	la, lb := 0, c
	fmt.Fprintf(out, "0 ")
	for i := 1; i < n; i++ {
		la, lb = min(la, lb)+a[i-1], min(la+c, lb)+b[i-1]
		fmt.Fprintf(out, "%d ", min(la, lb))
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1249E(in, os.Stdout)
}
