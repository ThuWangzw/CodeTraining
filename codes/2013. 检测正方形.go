type DetectSquares struct {
	nodes [1001][1001]int
}

func Constructor() DetectSquares {
	return DetectSquares{}
}

func (this *DetectSquares) Add(point []int) {
	this.nodes[point[0]][point[1]]++
}

func (this *DetectSquares) Count(point []int) int {
	count := 0
	for x := point[0] - 1; x >= 0; x-- {
		y := point[1] - (point[0] - x)
		if y < 0 {
			break
		}
		count += this.nodes[x][y] * this.nodes[x][point[1]] * this.nodes[point[0]][y]
	}
	for x := point[0] + 1; x <= 1000; x++ {
		y := point[1] + (x - point[0])
		if y > 1000 {
			break
		}
		count += this.nodes[x][y] * this.nodes[x][point[1]] * this.nodes[point[0]][y]
	}
	for x := point[0] - 1; x >= 0; x-- {
		y := point[1] + (point[0] - x)
		if y > 1000 {
			break
		}
		count += this.nodes[x][y] * this.nodes[x][point[1]] * this.nodes[point[0]][y]
	}
	for x := point[0] + 1; x <= 1000; x++ {
		y := point[1] - (x - point[0])
		if y < 0 {
			break
		}
		count += this.nodes[x][y] * this.nodes[x][point[1]] * this.nodes[point[0]][y]
	}
	return count
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */