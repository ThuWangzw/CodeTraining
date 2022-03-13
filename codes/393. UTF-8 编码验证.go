func validUtf8(data []int) bool {
	n := len(data)
	i := 0
	for i < n {
		// find first 0
		idx := 7
		for idx >= 0 {
			if ((^data[i]) & (1 << idx)) > 0 {
				break
			}
			idx--
		}
		if idx == 6 || idx < 3 {
			return false
		}
		if idx == 7 {
			i++
		} else {
			maskA := 0
			for j := 7; j > idx; j-- {
				maskA += 1 << j
			}
			if !(data[i]&maskA == maskA) {
				return false
			}
			cnt := 6 - idx
			maskB := 1 << 7
			for j := i + 1; j <= i+cnt; j++ {
				if j >= n {
					return false
				}
				if (data[j]&maskB != maskB) || ((^data[j])&(1<<6)) == 0 {
					return false
				}
			}
			i += cnt + 1
		}
	}
	return true
}