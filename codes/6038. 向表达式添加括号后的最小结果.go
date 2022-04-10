func minimizeResult(expression string) string {
	// find +
	plusIdx := -1
	for k, ch := range expression {
		if ch == '+' {
			plusIdx = k
			break
		}
	}
	left := expression[:plusIdx]
	right := expression[plusIdx+1:]
	ans := ""
	minans := math.MaxInt32
	for leftIdx := 0; leftIdx < len(left); leftIdx++ {
		leftleft := 1
		if leftIdx > 0 {
			leftleft, _ = strconv.Atoi(left[:leftIdx])
		}
		leftright, _ := strconv.Atoi(left[leftIdx:])
		for rightIdx := 0; rightIdx < len(right); rightIdx++ {
			rightleft, _ := strconv.Atoi(right[:rightIdx+1])
			rightright := 1
			if rightIdx < len(right)-1 {
				rightright, _ = strconv.Atoi(right[rightIdx+1:])
			}
			if minans > leftleft*(leftright+rightleft)*rightright {
				minans = leftleft * (leftright + rightleft) * rightright
				ans = left[:leftIdx] + "(" + left[leftIdx:] + "+" + right[:rightIdx+1] + ")" + right[rightIdx+1:]
			}
		}
	}
	return ans
}