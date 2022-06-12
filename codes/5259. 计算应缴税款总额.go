func bsearch(brackets [][]int, income int) int {
	begin := 0
	end := len(brackets) - 1
	for begin < end-3 {
		mid := (begin + end) / 2
		if brackets[mid][0] <= income {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	for i := begin; i <= end; i++ {
		if brackets[i][0] > income {
			return i
		}
	}
	return len(brackets) - 1
}

func calculateTax(brackets [][]int, income int) float64 {
	idx := bsearch(brackets, income)
	var tex float64
	for i := idx; i >= 0; i-- {
		if i > 0 {
			tex += float64(income-brackets[i-1][0]) * float64(brackets[i][1]) / 100
			income = brackets[i-1][0]
		} else {
			tex += float64(income) * float64(brackets[i][1]) / 100
		}
	}
	return tex
}