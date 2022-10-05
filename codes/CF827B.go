package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF827B(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	t := (n - 1) / k
	m := (n - 1) % k
	if m == 0 {
		fmt.Fprintln(out, t*2)
	} else if m == 1 {
		fmt.Fprintln(out, t*2+1)
	} else {
		fmt.Fprintln(out, t*2+2)
	}
	idx := 2
	for i := 0; i < m; i++ {
		for j := 0; j < t; j++ {
			fmt.Fprintln(out, idx, idx+1)
			idx += 1
		}
		fmt.Fprintln(out, idx, 1)
		idx += 1
	}
	for i := 0; i < k-m; i++ {
		for j := 0; j < t-1; j++ {
			fmt.Fprintln(out, idx, idx+1)
			idx += 1
		}
		fmt.Fprintln(out, idx, 1)
		idx += 1
	}
}

func main() {
	// read, _ := os.Open("input.txt")
	read := os.Stdin
	CF827B(read, os.Stdout)
}
