func fib(n int) int {
	last := -1
	cur := 1
	next := 0
	for i:=0; i<=n; i++ {
		next = last + cur
		last = cur
		cur = next
	}
	return cur
}