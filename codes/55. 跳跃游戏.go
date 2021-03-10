// 首先注意到如果下标为i的点可以到达，那么下标小于i的点均可以到达，因此只需要遍历一遍数组，维护一下最远可以达到的点即可
func canJump(nums []int) bool {
	max_idx := 0
	n := len(nums)
	for i := 0; i <= max_idx; i++ {
		max_idx = int(math.Max(float64(max_idx), float64(i+nums[i])))
		if max_idx >= n-1 {
			return true
		}
	}
	return false
}