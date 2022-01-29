type UF struct {
	parents []int
	sizes   []int
}

func Constructor(n int) *UF {
	sizes := make([]int, n)
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		sizes[i] = 1
		parents[i] = i
	}
	return &UF{
		parents: parents,
		sizes:   sizes,
	}
}

func (uf *UF) find(i int) int {
	for uf.parents[i] != i {
		uf.parents[i] = uf.parents[uf.parents[i]]
		i = uf.parents[i]
	}
	return i
}

func (uf *UF) union(i int, j int) int {
	ip, jp := uf.find(i), uf.find(j)
	if uf.sizes[ip] > uf.sizes[jp] {
		i, j, ip, jp = j, i, jp, ip
	}
	uf.parents[ip] = jp
	uf.sizes[jp] += uf.sizes[ip]
	return jp
}

func (uf *UF) connected(i int, j int) bool {
	return uf.find(i) == uf.find(j)
}

func qsort(connections [][]int, begin int, end int) {
	if begin >= end {
		return
	}
	pivot := connections[begin][2]
	left := begin + 1
	right := end
	for left <= right {
		if connections[left][2] <= pivot {
			left++
		} else {
			connections[left], connections[right] = connections[right], connections[left]
			right--
		}
	}
	connections[right], connections[begin] = connections[begin], connections[right]
	qsort(connections, begin, right-1)
	qsort(connections, right+1, end)
}

func minimumCost(n int, connections [][]int) int {
	uf := Constructor(n)
	qsort(connections, 0, len(connections)-1)
	cost := 0
	for _, edge := range connections {
		if !uf.connected(edge[0]-1, edge[1]-1) {
			cost += edge[2]
			uf.union(edge[0]-1, edge[1]-1)
		}
	}
	root := uf.find(0)
	for i := 0; i < n; i++ {
		if uf.find(i) != root {
			return -1
		}
	}
	return cost
}