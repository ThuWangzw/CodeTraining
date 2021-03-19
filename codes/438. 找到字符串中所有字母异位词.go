// 维护一个固定长度的滑动窗口
func index(c byte) int {
	return int(c - 'a')
}

func findAnagrams(s string, p string) []int {
	n := len(s)
	m := len(p)
	result := make([]int, 0)
	if n < m {
		return result
	}
	s_cnt := make([]int, 26, 26)
	p_cnt := make([]int, 26, 26)
	for i := 0; i < m; i++ {
		s_cnt[index(s[i])]++
		p_cnt[index(p[i])]++
	}
	diff := 0
	for i := 0; i < 26; i++ {
		if s_cnt[i] != p_cnt[i] {
			diff++
		}
	}
	if diff == 0 {
		result = append(result, 0)
	}
	left := 0
	right := m - 1
	for right < n-1 {
		right++
		left++
		s_cnt[index(s[right])]++
		s_cnt[index(s[left-1])]--
		right_cnt := s_cnt[index(s[right])]
		left_cnt := s_cnt[index(s[left-1])]
		p_right_cnt := p_cnt[index(s[right])]
		p_left_cnt := p_cnt[index(s[left-1])]
		if s[right] != s[left-1] {
			if right_cnt == p_right_cnt {
				diff--
			} else if right_cnt-1 == p_right_cnt {
				diff++
			}
			if left_cnt == p_left_cnt {
				diff--
			} else if left_cnt+1 == p_left_cnt {
				diff++
			}
		}
		if diff == 0 {
			result = append(result, left)
		}
	}

	return result
}