// 方法1：用数学的思想直接求

func countBits(num int) []int {
	result := make([]int, num+1, num+1)
	result[0] = 0
	end := 1
	for {
		i := 0
		for i = 0; (i < end) && (i+end < num+1); i++ {
			result[end+i] = result[i] + 1
		}
		if (i + end) == (num + 1) {
			break
		}
		end *= 2
	}
	return result
}

// 方法2： 使用动态规划的思想，该问题是存在很多种最优子结构的
func countBits(num int) []int {
	result := make([]int, num+1, num+1)
	result[0] = 0
	for i := 1; i < num+1; i++ {
		result[i] = result[i/2] + (i & 1)
	}
	return result
}