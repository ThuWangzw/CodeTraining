func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}
	list1Map := make(map[string]int)
	for i, name := range list1 {
		list1Map[name] = i
	}
	minNames := make([]string, 0)
	minCnt := math.MaxInt32
	for i, name := range list2 {
		if cnt1, ok := list1Map[name]; ok {
			if cnt1+i < minCnt {
				minNames = []string{name}
				minCnt = cnt1 + i
			} else if cnt1+i == minCnt {
				minNames = append(minNames, name)
			}
		}
	}
	return minNames
}