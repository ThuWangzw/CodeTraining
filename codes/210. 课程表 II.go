func findOrderIter(idx int, graph [][]int, ans *[]int, state []int) bool {
	if state[idx] == 1 {
		return false
	}
	if state[idx] == 2 {
		return true
	}
	state[idx] = 1
	for _, child := range graph[idx] {
		if !findOrderIter(child, graph, ans, state) {
			return false
		}
	}
	state[idx] = 2
	*ans = append(*ans, idx)
	return true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}
	notstart := make([]bool, numCourses)
	for _, link := range prerequisites {
		graph[link[1]] = append(graph[link[1]], link[0])
		notstart[link[0]] = true
	}
	starts := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if !notstart[i] {
			starts = append(starts, i)
		}
	}
	ans := make([]int, 0)
	state := make([]int, numCourses)
	for _, start := range starts {
		tmpans := make([]int, 0)
		if !findOrderIter(start, graph, &tmpans, state) {
			return make([]int, 0)
		}
		ans = append(ans, tmpans...)
	}
	if len(ans) != numCourses {
		return make([]int, 0)
	}
	for i := 0; i <= (numCourses-1)/2; i++ {
		ans[i], ans[numCourses-1-i] = ans[numCourses-1-i], ans[i]
	}
	return ans
}