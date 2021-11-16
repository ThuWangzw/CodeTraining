// 脑筋急转弯，计算节点的数目
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

func isRectangleCover(rectangles [][]int) bool {
	points := make(map[string]int)
	minx := math.MaxInt32
	miny := math.MaxInt32
	maxy := math.MinInt32
	maxx := math.MinInt32
	for _, rectangle := range rectangles {
		minx = min(minx, rectangle[0])
		maxx = max(maxx, rectangle[2])
		miny = min(miny, rectangle[1])
		maxy = max(maxy, rectangle[3])
		points[strconv.Itoa(rectangle[0])+","+strconv.Itoa(rectangle[1])]++
		points[strconv.Itoa(rectangle[0])+","+strconv.Itoa(rectangle[3])]++
		points[strconv.Itoa(rectangle[2])+","+strconv.Itoa(rectangle[1])]++
		points[strconv.Itoa(rectangle[2])+","+strconv.Itoa(rectangle[3])]++
	}
	corners := []string{
		strconv.Itoa(minx) + "," + strconv.Itoa(miny),
		strconv.Itoa(minx) + "," + strconv.Itoa(maxy),
		strconv.Itoa(maxx) + "," + strconv.Itoa(miny),
		strconv.Itoa(maxx) + "," + strconv.Itoa(maxy),
	}
	for _, corner := range corners {
		if cnt, ok := points[corner]; !ok || cnt != 1 {
			return false
		}
		delete(points, corner)
	}
	for _, cnt := range points {
		if cnt != 2 && cnt != 4 {
			return false
		}
	}
	return true
}