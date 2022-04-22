func sort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func largestInteger(num int) int {
	numsFlag := make([]bool, 0)
	numsa := make([]int, 0)
	numsb := make([]int, 0)
	for num > 0 {
		bit := num % 10
		numsFlag = append(numsFlag, bit%2 == 0)
		if bit%2 == 0 {
			numsa = append(numsa, bit)
		} else {
			numsb = append(numsb, bit)
		}
		num /= 10
	}
	sort(numsa)
	sort(numsb)
	i := 0
	j := 0
	newnum := 0
	for k := len(numsFlag) - 1; k >= 0; k-- {
		if numsFlag[k] {
			newnum = newnum*10 + numsa[i]
			i++
		} else {
			newnum = newnum*10 + numsb[j]
			j++
		}
	}
	return newnum
}