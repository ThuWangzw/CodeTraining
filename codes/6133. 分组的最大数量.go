func maximumGroups(grades []int) int {
	lastcnt := 0
	lastsum := 0
	cnt := 0
	sum := 0
	ans := 0
	sort.Ints(grades)
	for _, grade := range grades {
		cnt++
		sum += grade
		if cnt > lastcnt && sum > lastsum {
			ans++
			lastcnt, lastsum, cnt, sum = cnt, sum, 0, 0
		}
	}
	if cnt > lastcnt && sum > lastsum {
		ans++
	}
	return ans
}