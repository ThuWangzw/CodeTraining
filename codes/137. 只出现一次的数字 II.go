// 可以构建自动机来做，并不是天马行空的想法
func singleNumber(nums []int) int {
	one := 0
	two := 0
	for _, num := range nums {
		one = (^two) & (one ^ num)
		two = (^one) & (two ^ num)
	}
	return one
}