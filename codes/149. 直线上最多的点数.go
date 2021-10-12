// 哈希表
func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Line [2]float32

func maxPoints(points [][]int) int {
	lineMap := make(map[Line]map[[2]int]bool)
	n := len(points)
	if n == 0 {
		return 0
	}
	maxlen := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			u := points[i]
			v := points[j]
			var line Line
			if u[0] == v[0] {
				line = Line([2]float32{math.MaxFloat32, float32(u[0])})
			} else {
				k := float32(u[1]-v[1]) / float32(u[0]-v[0])
				b := float32(u[0]*v[1]-u[1]*v[0]) / float32(u[0]-v[0])
				line = Line([2]float32{k, b})
			}
			_, ok := lineMap[line]
			if !ok {
				lineMap[line] = make(map[[2]int]bool)
			}
			lineMap[line][[2]int{u[0], u[1]}] = true
			lineMap[line][[2]int{v[0], v[1]}] = true
			maxlen = max(maxlen, len(lineMap[line]))
		}
	}
	return maxlen
}