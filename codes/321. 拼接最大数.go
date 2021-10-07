// 可以用化简的思想，先想一下如果只有一个数组的做法（简单的单调栈），再遍历k得到两个数组的解法
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m := len(nums1)
	n := len(nums2)
	ans := []int{0}
	for i := max(0, k-n); i <= min(k, m); i++ {
		a := findMax(nums1, i)
		b := findMax(nums2, k-i)
		res := merge(a, b)
		if less(ans, res) {
			ans = res
		}
	}
	return ans
}

func less(a []int, b []int) bool {
	m := len(a)
	n := len(b)
	if m == 0 && n > 0 {
		return true
	}
	end := min(m, n)
	for i := 0; i < end; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	if m >= n {
		return false
	}
	return true
}

func findMax(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0, k)
	if k == 0 {
		return ans
	} else if k == n {
		ans = append(ans, nums...)
		return ans
	}
	removed := 0
	for i := 0; i < n; i++ {
		num := nums[i]
		if len(ans) > 0 {
			top := ans[len(ans)-1]
			for removed < n-k && top < num {
				removed++
				ans = ans[:len(ans)-1]
				if len(ans) == 0 {
					break
				}
				top = ans[len(ans)-1]
			}
		}
		ans = append(ans, num)
	}
	if removed < n-k {
		ans = ans[:len(ans)-(n-k-removed)]
	}
	return ans
}

func merge(a []int, b []int) []int {
	m := len(a)
	n := len(b)
	ans := make([]int, m+n)
	i := 0
	j := 0
	for i < m && j < n {
		if a[i] < b[j] || (a[i] == b[j] && less(a[i:], b[j:])) {
			ans[i+j] = b[j]
			j++
		} else {
			ans[i+j] = a[i]
			i++
		}
	}
	for i < m {
		ans[i+j] = a[i]
		i++
	}
	for j < n {
		ans[i+j] = b[j]
		j++
	}
	return ans
}