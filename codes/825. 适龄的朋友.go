func numFriendRequests(ages []int) int {
	ageCount := make([]int, 121)
	for _, age := range ages {
		ageCount[age]++
	}
	ans := 0
	for age, count := range ageCount {
		if count <= 0 {
			continue
		}
		for yage := int(0.5*float32(age)+7) + 1; yage < age; yage++ {
			ans += ageCount[yage] * count
		}
		if age > int(0.5*float32(age)+7) {
			ans += count * (count - 1)
		}
	}
	return ans
}