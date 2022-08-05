package main

import (
	"bufio"
	"fmt"
	"os"
)

func CF808D(nums []int) string {
	numMap := make(map[int64]int)
	var sum int64 = 0
	for _, num := range nums {
		numMap[int64(num)]++
		sum += int64(num)
	}
	if sum%2 != 0 {
		return "NO"
	}

	var prefix int64 = 0
	leftMap := make(map[int64]int)
	for _, num := range nums {
		prefix += int64(num)
		right := sum - prefix
		if prefix == right {
			return "YES"
		}
		leftMap[int64(num)]++
		numMap[int64(num)]--
		if prefix > right && leftMap[(prefix-right)/2] > 0 {
			return "YES"
		} else if right > prefix && numMap[(right-prefix)/2] > 0 {
			return "YES"
		}
	}
	return "NO"
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
	fmt.Fprintln(out, CF808D(as))
}
