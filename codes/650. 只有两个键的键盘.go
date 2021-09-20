// 关键是发现复制+粘贴与整数除法等价
// 方法一使用动态规划比较直接
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func minSteps(n int) int {
	counts := make([]int, n+1)
	for i := 2; i <= n; i++ {
		counts[i] = i
		for j := 1; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				counts[i] = min(counts[i], i/j+counts[j])
				counts[i] = min(counts[i], j+counts[i/j])
			}
		}
	}
	return counts[n]
}

// 方法二发现等价于质因数分解
func minSteps(n int) int {
	count := 0
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			count += i
		}
	}
	if n > 1 {
		count += n
	}
	return count
}