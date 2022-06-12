func bs(target int64, success int64, potions []int) int {
	begin := 0
	end := len(potions) - 1
	for begin < end-3 {
		mid := (begin + end) / 2
		if target*int64(potions[mid]) < success {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	for i := begin; i <= end; i++ {
		if target*int64(potions[i]) >= success {
			return len(potions) - i
		}
	}
	return 0
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		ans[i] = bs(int64(spells[i]), success, potions)
	}
	return ans
}