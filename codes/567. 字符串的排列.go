// 维护一个长度固定的窗口
func index(c byte) int {
	return int(c-'a')
}

func checkInclusion(s1 string, s2 string) bool {
	t_cnt := make([]int, 26, 26)
	// count alphabet in t
	for _, c := range s1 {
		t_cnt[index(byte(c))]++
	}
	left := 0
	right := 0
	n := len(s2)
	s_cnt := make([]int, 26, 26)
	for right < n {
		s_cnt[index(s2[right])]++
		for {
			less_or_large := true
			is_equal := true
			for i := 0; i < 26; i++ {
				if s_cnt[i] != t_cnt[i] {
					is_equal = false
					if s_cnt[i] < t_cnt[i] {
						less_or_large = false
					}
				}
			}
			if is_equal {
				return true
			} else if less_or_large {
				s_cnt[index(s2[left])]--
				left++
			} else if !less_or_large && (left>0) {
				left--
				s_cnt[index(s2[left])]++
				break
			} else {
				break
			}
		}
		right++
	}
	return false
}