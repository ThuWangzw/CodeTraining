func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func index(c byte) int {
	return int(c)
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	cnt := make([]int, 128, 128)
	left := 0
	right := 0
	max_len := 0
	for right < n {
		cnt[index(s[right])]++
		ok := true
		for _, v := range cnt {
			if v > 1 {
				ok = false
				break
			}
		}
		for !ok {
			cnt[index(s[left])]--
			left++
			ok = true
			for _, v := range cnt {
				if v > 1 {
					ok = false
					break
				}
			}
		}
		right++
		max_len = max(max_len, right-left)
	}
	return max_len
}