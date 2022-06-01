func isNumber(s string) bool {
	// skip space
	begin := len(s)
	for i, ch := range s {
		if ch != ' ' {
			begin = i
			break
		}
	}
	// get float
	flag, floatlen := isFloat(s[begin:])
	if !flag {
		return false
	}
	begin += floatlen
	if begin < len(s) && (s[begin] == 'e' || s[begin] == 'E') {
		begin++
		first := begin
		if first < len(s) && (s[first] == '+' || s[first] == '-') {
			begin++
			first++
		}
		for first < len(s) && s[first] >= '0' && s[first] <= '9' {
			first++
		}
		if first == begin {
			return false
		}
		begin = first
	}
	// skip space
	for begin < len(s) && s[begin] == ' ' {
		begin++
	}
	return begin == len(s)
}

func isFloat(s string) (bool, int) {
	n := len(s)
	if n == 0 {
		return false, 0
	}
	begin := 0
	// positive or negative
	if s[0] == '+' || s[0] == '-' {
		begin++
	}
	// first int
	first := begin
	for first < n && s[first] >= '0' && s[first] <= '9' {
		first++
	}
	if first != begin && (first == n || s[first] != '.') {
		return true, first
	}
	if first == begin {
		if first == n || s[first] != '.' {
			return false, first
		} else {
			first++
			begin++
			for first < n && s[first] >= '0' && s[first] <= '9' {
				first++
			}
			if first == begin {
				return false, first
			} else {
				return true, first
			}
		}
	} else {
		first++
		begin = first
		for first < n && s[first] >= '0' && s[first] <= '9' {
			first++
		}
		return true, first
	}
}