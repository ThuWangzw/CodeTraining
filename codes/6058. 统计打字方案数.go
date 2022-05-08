func countChar(s, k int, buff map[int]int) int {
	if cnt, ok := buff[s]; ok {
		return cnt
	}
	if s == 1 {
		buff[1] = 1
		return 1
	} else if s == 2 {
		buff[2] = 2
		return 2
	} else if s == 3 {
		buff[3] = 4
		return 4
	} else if s == 0 {
		buff[0] = 1
		return 1
	}
	for i := 1; i <= k; i++ {
		buff[s] = (buff[s] + countChar(s-i, k, buff)) % (1e9 + 7)
	}
	return buff[s]
}

func countTexts(pressedKeys string) int {
	left := 0
	right := 0
	n := len(pressedKeys)
	var ans int64
	ans = 1
	keyCount := []int{0, 0, 3, 3, 3, 3, 3, 4, 3, 4}
	buff3 := make(map[int]int)
	buff4 := make(map[int]int)
	for right < n {
		if pressedKeys[left] == pressedKeys[right] {
			right++
		} else {
			k := keyCount[int(pressedKeys[left]-'0')]
			if k == 3 {
				ans = (ans * int64(countChar(right-left, 3, buff3))) % (1e9 + 7)
			} else {
				ans = (ans * int64(countChar(right-left, 4, buff4))) % (1e9 + 7)
			}
			left = right
		}
	}
	k := keyCount[int(pressedKeys[left]-'0')]
	if k == 3 {
		ans = (ans * int64(countChar(right-left, 3, buff3))) % (1e9 + 7)
	} else {
		ans = (ans * int64(countChar(right-left, 4, buff4))) % (1e9 + 7)
	}
	return int(ans)
}