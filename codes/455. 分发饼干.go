// 贪心算法
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	cnt := 0
	m := len(g)
	n := len(s)
	i := m - 1
	for j := n - 1; j >= 0 && i >= 0; j-- {
		if s[j] >= g[i] {
			i--
			cnt++
		} else {
			i--
			j++
		}
	}
	return cnt
}