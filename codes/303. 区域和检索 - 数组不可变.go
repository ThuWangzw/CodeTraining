// 经典的前缀和
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sum := 0
	sums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sums[i] = sum
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	} else {
		return this.sums[right] - this.sums[left-1]
	}
}