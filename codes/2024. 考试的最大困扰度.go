// 滑动窗口，处理区间的计数问题
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxConsecutiveAnswersOfChar(answerKey string, k int, c byte) int {
	n := len(answerKey)
	var left, right, sum, maxLen int
	for ; right < n; right++ {
		if answerKey[right] != c {
			sum++
		}
		for sum > k {
			if answerKey[left] != c {
				sum--
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswersOfChar(answerKey, k, 'T'), maxConsecutiveAnswersOfChar(answerKey, k, 'F'))
}