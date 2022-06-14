func mul(a int32, b int32) (int32, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}
	half := b / 2
	halfans, overflow := mul(a, half)
	if overflow {
		return math.MaxInt32, true
	}
	ans := halfans + halfans
	if (halfans > 0 && ans < halfans) || (halfans < 0 && ans > halfans) {
		return math.MaxInt32, true
	}
	if b%2 != 0 {
		tmp := ans + a
		if (a > 0 && tmp < ans) || (a < 0 && tmp > ans) {
			return math.MaxInt32, true
		}
		ans = tmp
	}
	return ans, false
}

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if a == math.MinInt32 && b == 1 {
		return math.MinInt32
	}

	if a == math.MinInt32 && b == math.MinInt32 {
		return 1
	}
	if b == math.MinInt32 {
		return 0
	}
	if a == math.MinInt32 {

	}
	flag := 1
	if a == math.MinInt32 {
		if b > 0 {
			b = -b
			flag = -1
		}
	} else if a > 0 && b < 0 {
		flag = -1
		b = -b
	} else if b > 0 && a < 0 {
		flag = -1
		a = -a
	} else if a < 0 && b < 0 {
		a = -a
		b = -b
	}

	var left int32 = 0
	var right int32 = int32(a)
	if a == math.MinInt32 {
		right = math.MaxInt32
	}
	for left < right-3 {
		mid := (right-left)/2 + left
		midans, _ := mul(int32(b), int32(mid))
		if midans < int32(a) || (int32(a) == math.MinInt32 && midans < 0) {
			left = mid
		} else {
			right = mid
		}
	}
	for i := left; i <= right; i++ {
		ans, overflow := mul(int32(b), i)
		if overflow || (ans > int32(a) && a != math.MinInt32) {
			return flag * int(i-1)
		}
	}
	return flag * int(right)
}