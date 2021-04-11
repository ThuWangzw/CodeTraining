// 使用一趟扫描即可，遇到新的值就“插入到”合适的位置
func sortColors(nums []int) {
	end := make([]int, 3)
	n := len(nums)
	for i := 0; i < n; i++ {
		switch nums[i] {
		case 0:
			nums[end[2]], nums[end[1]], nums[end[0]] = nums[end[1]], nums[end[0]], 0
			end[0]++
			end[1]++
			end[2]++
		case 1:
			nums[end[2]], nums[end[1]] = nums[end[1]], 1
			end[1]++
			end[2]++
		case 2:
			end[2]++
		}
	}
}