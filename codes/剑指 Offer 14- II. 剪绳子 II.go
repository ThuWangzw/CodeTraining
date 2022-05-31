func pow3(n int64) int64 {
	switch n {
	case 0:
		return 1
	case 1:
		return 3
	case 2:
		return 9
	default:
		if n%2 == 1 {
			return (pow3(n/2) * pow3(n/2) * 3) % (1e9 + 7)
		} else {
			return (pow3(n/2) * pow3(n/2)) % (1e9 + 7)
		}
	}
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	cnt3 := int64(n / 3)
	cnt2 := int64(n % 3)
	if cnt2 == 0 {
		return int(pow3(cnt3))
	} else if cnt2 == 1 {
		return int((pow3(cnt3-1) * 4) % (1e9 + 7))
	} else {
		return int((pow3(cnt3) * 2) % (1e9 + 7))
	}
}