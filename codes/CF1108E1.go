package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func MostChosen(rawnums []int, intervals [][]int, maxi int) (delta int, chosen []int) {
	nums := make([]int, 0, len(rawnums))
	nums = append(nums, rawnums...)

	chosen = make([]int, 0)
	for j, interval := range intervals {
		l, r := interval[0], interval[1]
		if maxi < l-1 || maxi > r-1 {
			chosen = append(chosen, j+1)
			nums[l-1]--
			if r < len(nums) {
				nums[r]++
			}
		}
	}
	minnum := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
		if minnum > nums[i] {
			minnum = nums[i]
		}
	}
	delta = nums[maxi] - minnum
	return delta, chosen
}

func CF1108E1(read io.Reader, write io.Writer) {
	in := bufio.NewReader(read)
	out := bufio.NewWriter(write)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}
	for i := n - 1; i > 0; i-- {
		nums[i] -= nums[i-1]
	}

	intervals := make([][]int, m)
	for j := 0; j < m; j++ {
		intervals[j] = make([]int, 2)
		fmt.Fscan(in, &intervals[j][0], &intervals[j][1])
	}

	maxdelta, maxchosen := 0, make([]int, 0)
	for i := 0; i < n; i++ {
		delta, chosen := MostChosen(nums, intervals, i)
		if maxdelta < delta {
			maxdelta = delta
			maxchosen = chosen
		}
	}

	fmt.Fprintln(out, maxdelta)
	fmt.Fprintln(out, len(maxchosen))
	for _, num := range maxchosen {
		fmt.Fprintf(out, "%d ", num)
	}
}

func main() {
	// in, _ := os.Open("input.txt")
	in := os.Stdin
	CF1108E1(in, os.Stdout)
}
