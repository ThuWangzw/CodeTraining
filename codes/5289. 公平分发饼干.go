var maxcookie int

func cookie2children(cookies []int, children []int) {
	if len(cookies) == 0 {
		maxthis := 0
		for _, child := range children {
			maxthis = max(maxthis, child)
		}
		maxcookie = min(maxcookie, maxthis)
		return
	}
	for i := 0; i < len(children); i++ {
		children[i] += cookies[0]
		cookie2children(cookies[1:], children)
		children[i] -= cookies[0]
	}
}

func distributeCookies(cookies []int, k int) int {
	maxcookie = math.MaxInt32
	children := make([]int, k)
	cookie2children(cookies, children)
	return maxcookie
}

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