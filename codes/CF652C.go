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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CF652C(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	numIndex := make(map[int]int)
	nums := make([]int, n)
	minLast := make(map[int]int)

	for i := 0; i < n; i++ {
		num := 0
		fmt.Fscan(in, &num)
		numIndex[num-1] = i
		minLast[num-1] = -1
		nums[i] = num - 1
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if numIndex[b] < numIndex[a] {
			minLast[a] = max(minLast[a], numIndex[b])
		}
		if numIndex[a] < numIndex[b] {
			minLast[b] = max(minLast[b], numIndex[a])
		}
	}
	var ans int64
	last := -1
	for _, num := range nums {
		last = max(last, minLast[num])
		ans += int64(numIndex[num] - last)
	}
	fmt.Fprintln(out, ans)
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF652C(in, os.Stdout)
}
