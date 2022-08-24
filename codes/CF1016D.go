package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1016D(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var m, n int
	fmt.Fscan(in, &m, &n)
	var xra, xrb int
	as := make([]int, m)
	bs := make([]int, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &as[i])
		xra ^= as[i]
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &bs[i])
		xrb ^= bs[i]
	}
	if xra != xrb {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Fprintf(out, "%d ", xra^as[0]^bs[0])
		} else if i < n-1 {
			fmt.Fprintf(out, "%d ", bs[i])
		} else {
			fmt.Fprintf(out, "%d\n", bs[i])
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				fmt.Fprintf(out, "%d ", as[i])
			} else if j < n-1 {
				fmt.Fprintf(out, "0 ")
			} else {
				fmt.Fprintf(out, "0\n")
			}
		}

	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1016D(in, os.Stdout)
}
