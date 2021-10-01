// 投票算法，找到数组中数量超过一半的鼠
func majorityElement(nums []int) int {
	candidate := 0
	cnt := 0
	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}
		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}