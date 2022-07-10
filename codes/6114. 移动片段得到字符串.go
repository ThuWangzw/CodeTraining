func canChange(start string, target string) bool {
	n := len(start)
	i := 0
	for j, ch := range target {
		switch ch {
		case '_':
			continue
		case 'L':
			find := false
			for i < n {
				if start[i] == '_' {
					i++
				} else if start[i] == 'L' {
					if i < j {
						return false
					}
					i++
					find = true
					break
				} else {
					return false
				}
			}
			if !find {
				return false
			}
		case 'R':
			find := false
			for i < n {
				if start[i] == '_' {
					i++
				} else if start[i] == 'R' {
					if i > j {
						return false
					}
					i++
					find = true
					break
				} else {
					return false
				}
			}
			if !find {
				return false
			}
		}
	}
	for i < n {
		if start[i] != '_' {
			return false
		}
		i++
	}
	return true
}