var nBit1Count map[int]int = make(map[int]int)

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	if n < 10 {
		return 1
	}
	return countDigitOneIter(num2ary(n))
}

func countDigitOneIter(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 0
		} else {
			return 1
		}
	}
	ans := nums[len(nums)-1] * count1inN(len(nums)-1)
	if nums[len(nums)-1] > 1 {
		ans += int(math.Pow(10, float64(len(nums)-1)))
	}
	ans += countDigitOneIter(nums[:len(nums)-1])
	if nums[len(nums)-1] == 1 {
		ans += ary2num(nums[:len(nums)-1]) + 1
	}
	return ans
}

// count 1 in N-bit number
func count1inN(bitCount int) int {
	if bitCount == 0 {
		return 0
	}
	if nBit1Count[bitCount] > 0 {
		return nBit1Count[bitCount]
	}
	nBit1Count[bitCount] = 10*count1inN(bitCount-1) + int(math.Pow(10, float64(bitCount-1)))
	return nBit1Count[bitCount]
}

func ary2num(nums []int) int {
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		ans = ans*10 + nums[i]
	}
	return ans
}

func num2ary(num int) []int {
	ans := make([]int, 0)
	for num > 0 {
		ans = append(ans, num%10)
		num /= 10
	}
	if len(ans) == 0 {
		ans = append(ans, 0)
	}
	return ans
}