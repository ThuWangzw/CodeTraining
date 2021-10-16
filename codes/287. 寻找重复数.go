// 这道题给数组型题目提供了两个思路，一个是可以看成链表，另一个是可以用计数法
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}