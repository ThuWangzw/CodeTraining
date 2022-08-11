package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CF1217C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		cnt := 0
		last := 0
		margin := int(math.Log2(float64(len(s))) + 2)
		for i, c := range s {
			if c == '0' {
				last++
			} else {
				sum := 0
				for j := i; j < min(len(s), i+margin); j++ {
					sum = sum*2 + int(s[j]-'0')
					if sum >= j-i+1 && j-i+1+last >= sum {
						cnt++
					}
				}

				last = 0
			}
		}
		fmt.Fprintln(out, cnt)
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1217C(in, os.Stdout)
}
