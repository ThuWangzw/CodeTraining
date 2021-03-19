// 滑动窗口法
func index(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a')
	} else {
		return int(c-'A')+26
	}
}

func minWindow(s string, t string) string {
	t_cnt := make([]int, 52, 52)
	// count alphabet in t
	for _, c := range t {
		t_cnt[index(byte(c))]++
	}
	left := 0
	right := 0
	min_len := math.MaxInt32
	min_left := -1
	min_right := -1
	n := len(s)
	s_cnt := make([]int, 52, 52)
	for right < n {
		s_cnt[index(s[right])]++
		for {
			contain := true
			for i := 0; i < 52; i++ {
				if s_cnt[i] < t_cnt[i] {
					contain = false
					break
				}
			}
			if contain && left < right {
				s_cnt[index(s[left])]--
				left++
			} else if !contain && (left > 0) {
				left--
				s_cnt[index(s[left])]++
				if min_len > right-left+1 {
					min_len = right - left + 1
					min_left = left
					min_right = right
				}

				break
			} else if contain {
				min_len = right - left + 1
				min_left = left
				min_right = right
				break
			} else {
				break
			}
		}
		right++
	}
	if min_len == math.MaxInt32 {
		return ""
	} else {
		return s[min_left : min_right+1]
	}
}