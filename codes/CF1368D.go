package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1368D(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n int
	cnts := make([]int, 20)
	for fmt.Fscan(in, &n); n > 0; n-- {
		var num int
		fmt.Fscan(in, &num)
		for i := 0; i < 20; i++ {
			if num&1 > 0 {
				cnts[i]++
			}
			num = num >> 1
		}
	}

	var ans int64
	for true {
		num := 0
		for i := 0; i < 20; i++ {
			if cnts[i] > 0 {
				cnts[i]--
				num |= 1 << i
			}
		}
		if num == 0 {
			break
		}
		ans += int64(num) * int64(num)
	}
	fmt.Fprintln(out, ans)
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1368D(in, os.Stdout)
}
