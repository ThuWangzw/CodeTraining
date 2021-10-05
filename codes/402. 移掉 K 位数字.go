// 单调栈： 首先意识到，移出一位数字最优解是移除掉a[i]>a[i+1]的那一位，然后推理得到每次都移出这一位，即可得到最优解，进而使用单调栈求解
func removeKdigits(num string, k int) string {
	ms := make([]byte, 0)
	n := len(num)
	i := 0
	cnt := 0
	for i = 0; i < n; i++ {
		for len(ms) > 0 && ms[len(ms)-1] > num[i] && cnt < k {
			ms = ms[:len(ms)-1]
			cnt++
		}
		ms = append(ms, num[i])
		if cnt == k {
			i++
			break
		}
	}
	ans := string(ms) + num[i:]
	if cnt < k {
		ans = ans[:len(ans)-(k-cnt)]
	}
	if ans == "" {
		ans = "0"
	}
	idx := 0
	for idx = 0; (idx < len(ans)-1) && (ans[idx] == '0'); idx++ {
	}
	return ans[idx:]
}