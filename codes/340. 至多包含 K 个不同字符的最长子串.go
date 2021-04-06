// 典型的双指针（滑动窗口）问题
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	i := 0
	j := -1
	n := len(s)
	charCnts := make([]int, 128)
	diffChars := 0
	maxLen := 0
	for j < n {
		// add j
		for diffChars <= k {
			maxLen = max(maxLen, j-i+1)
			j++
			if j == n {
				break
			}
			charCnts[int(s[j])]++
			if charCnts[int(s[j])] == 1 {
				diffChars++
			}
		}
		if j == n {
			maxLen = max(maxLen, j-i)
			break
		}
		//add i
		for diffChars > k {
			charCnts[int(s[i])]--
			if charCnts[int(s[i])] == 0 {
				diffChars--
			}
			i++
		}
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}