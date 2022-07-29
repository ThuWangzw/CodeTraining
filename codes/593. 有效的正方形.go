func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	dis2 := dis(p1, p2)
	dis3 := dis(p1, p3)
	dis4 := dis(p1, p4)
	if dis2 == dis3 && dis2*2 == dis4 {
		return checkdis(p2, p1, p4, p3) && checkdis(p3, p1, p4, p2)
	} else if dis2 == dis4 && dis2*2 == dis3 {
		return checkdis(p2, p1, p3, p4) && checkdis(p4, p1, p3, p2)
	} else if dis3 == dis4 && dis2 == dis4*2 {
		return checkdis(p3, p1, p2, p4) && checkdis(p4, p1, p2, p3)
	}
	return false
}

func checkdis(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	dis2 := dis(p1, p2)
	dis3 := dis(p1, p3)
	dis4 := dis(p1, p4)
	if dis2 == 0 {
		return false
	}
	return dis2 == dis3 && dis2*2 == dis4
}

func dis(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}