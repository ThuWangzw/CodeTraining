package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1066E(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	mod := 998244353
	var n, m int
	fmt.Fscan(in, &n, &m)
	var as, bs string
	fmt.Fscan(in, &as, &bs)

	presum := make([]int, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			presum[i] = int(bs[i] - '0')
		} else {
			presum[i] = presum[i-1] + int(bs[i]-'0')
		}
	}
	ans := make([]int, 0)
	last := 0
	for i := m - 1; i >= 0; i-- {
		base := last
		if i-m+n >= 0 && as[i-m+n] == '1' {
			base += presum[i]
		}
		ans = append(ans, base%2)
		last = base / 2
	}
	ans = append(ans, last)
	var res int64
	for i := len(ans) - 1; i >= 0; i-- {
		res = (res*2 + int64(ans[i])) % int64(mod)
	}
	fmt.Fprintln(out, res)
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1066E(in, os.Stdout)
}
