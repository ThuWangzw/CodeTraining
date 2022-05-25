func singleNumber(nums []int) int {
	one := 0
	two := 0
	for _, num := range nums {
		one, two = (^(num ^ two))&(num^one), (^one)&(num^two)
	}
	return two
}