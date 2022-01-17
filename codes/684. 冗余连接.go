type DisjointSet []int

func Constructor(n int) *DisjointSet {
	ds := make([]int, n)
	for i := 0; i < n; i++ {
		ds[i] = i
	}
	return (*DisjointSet)(&ds)
}

func (ds *DisjointSet) Union(i, j int) {
	ia := ds.Find(i)
	(*ds)[ia] = (*ds)[j]
}

func (ds *DisjointSet) Find(i int) int {
	if (*ds)[i] == i {
		return i
	} else {
		(*ds)[i] = ds.Find((*ds)[i])
		return (*ds)[i]
	}
}

func findRedundantConnection(edges [][]int) []int {
	var ans []int
	n := len(edges)
	ds := Constructor(n)
	for _, edge := range edges {
		if ds.Find(edge[0]-1) == ds.Find(edge[1]-1) {
			ans = edge
		}
		ds.Union(edge[0]-1, edge[1]-1)
	}
	return ans
}