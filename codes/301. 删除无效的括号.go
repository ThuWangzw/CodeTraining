// 递归
func removeIter(s string, idx int, lcnt int, rcnt int, rl int, rr int, cur string, ans map[string]bool) {
	n := len(s)
	if rl == 0 && rr == 0 && idx == n {
		cur += s[idx:]
		ans[cur] = true
		return
	}
	if idx == n {
		return
	}

	if s[idx] == '(' {
		if rl > 0 {
			removeIter(s, idx+1, lcnt, rcnt, rl-1, rr, cur, ans)
		}
		removeIter(s, idx+1, lcnt+1, rcnt, rl, rr, cur+"(", ans)
	} else if s[idx] == ')' {
		if rr > 0 {
			removeIter(s, idx+1, lcnt, rcnt, rl, rr-1, cur, ans)
		}
		if lcnt >= rcnt+1 {
			removeIter(s, idx+1, lcnt, rcnt+1, rl, rr, cur+")", ans)
		}
	} else {
		removeIter(s, idx+1, lcnt, rcnt, rl, rr, cur+string(s[idx]), ans)
	}
}

func removeInvalidParentheses(s string) []string {
	rlcnt := 0
	rrcnt := 0
	for _, ch := range s {
		switch ch {
		case '(':
			rlcnt++
		case ')':
			if rlcnt == 0 {
				rrcnt++
			} else {
				rlcnt--
			}
		}
	}
	ans := make([]string, 0)
	ansmap := make(map[string]bool)
	removeIter(s, 0, 0, 0, rlcnt, rrcnt, "", ansmap)
	for key, _ := range ansmap {
		ans = append(ans, key)
	}
	return ans
}