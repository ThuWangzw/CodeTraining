func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	ans := passengers[0] - 1
	i := 0
	for _, leave := range buses {
		begin := i
		for {
			if i-begin+1 <= capacity && i < len(passengers) && passengers[i] <= leave {
				i++
			} else {
				// end := i
				// if i-begin+1 > capacity {
				// 	end--
				// }
				for j := begin; j < i; j++ {
					if j == 0 || passengers[j-1] != passengers[j]-1 {
						ans = passengers[j] - 1
					}
				}
				if i-begin+1 <= capacity {
					if (i == len(passengers) || passengers[i] > leave) && (i == 0 || passengers[i-1] != leave) {
						ans = leave
					}
				}

				break
			}
		}
	}
	return ans
}