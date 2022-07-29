package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type GeoHorse [][]int64

func (gh GeoHorse) Len() int {
	return len(gh)
}

func (gh GeoHorse) Swap(a, b int) {
	gh[a], gh[b] = gh[b], gh[a]
}

func (gh GeoHorse) Less(a, b int) bool {
	return gh[a][1] < gh[b][1]
}

func CF175C(nums [][]int64, counts []int64) int64 {
	n := len(nums)
	sort.Sort(GeoHorse(nums))
	i := 0
	var score, last int64
	for k, cnt := range counts {
		cnt, last = cnt-last, cnt
		for i < n {
			if nums[i][0] <= cnt {
				score += int64(nums[i][0]) * int64(nums[i][1]) * int64(k+1)
				cnt -= nums[i][0]
				i++
			} else {
				nums[i][0] -= cnt
				score += int64(nums[i][1]) * int64(cnt) * int64(k+1)
				break
			}
		}
		k++
	}
	for i < n {
		score += int64(nums[i][1]) * int64(nums[i][0]) * int64(len(counts)+1)
		i++
	}
	return score
}

func main() {
	// read, _ := os.Open("input.txt")
	read := os.Stdin
	write := os.Stdout
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()
	var n, t int
	fmt.Fscan(in, &n)
	nums := make([][]int64, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int64, 2)
		fmt.Fscan(in, &nums[i][0], &nums[i][1])
	}
	fmt.Fscan(in, &t)
	counts := make([]int64, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &counts[i])
	}
	fmt.Fprintln(out, CF175C(nums, counts))
}
