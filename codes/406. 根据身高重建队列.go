type Peoples [][]int

func (ps Peoples) Len() int {
	return len(ps)
}
func (ps Peoples) Swap(i int, j int) {
	ps[j], ps[i] = ps[i], ps[j]
}
func (ps Peoples) Less(i int, j int) bool {
	return (ps[i][0] < ps[j][0]) || (ps[i][0] == ps[j][0] && ps[i][1] > ps[j][1])
}

func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	sort.Sort(Peoples(people))
	finalQueue := make([][]int, n)
	for _, p := range people {
		idx := p[1]
		j := 0
		i := 0
		for {
			if j == idx && len(finalQueue[i]) == 0 {
				finalQueue[i] = p
				break
			}
			if len(finalQueue[i]) == 0 {
				j++
			}
			i++
		}
	}
	return finalQueue
}