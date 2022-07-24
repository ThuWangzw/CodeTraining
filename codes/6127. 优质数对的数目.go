func countExcellentPairs(nums []int, k int) int64 {
	dupcnt := make(map[int]int)
	for _, num := range nums {
		dupcnt[num]++
	}
	cnts := make([]int, 0)
	for num, _ := range dupcnt {
		cnts = append(cnts, count1(num))
	}
	sort.Ints(cnts)
	var ans int64
	for _, num := range cnts {
		ans += int64(find(cnts, k-num))
	}

	return ans
}

func count1(num int) int {
	cnt := 0
	for num > 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}

func find(nums []int, v int) int {
	left := 0
	right := len(nums) - 1
	n := len(nums)
	for left < right-3 {
		mid := (right-left)/2 + left
		if nums[mid] >= v {
			right = mid
		} else {
			left = mid
		}
	}
	for left <= right {
		if nums[left] >= v {
			return n - left
		}
		left++
	}
	return 0
}