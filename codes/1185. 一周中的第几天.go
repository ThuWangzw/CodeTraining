func dayOfTheWeek(day int, month int, year int) string {
	daycnt := 0
	for i := 1971; i < year; i++ {
		if i%4 == 0 && (i%100 != 0 || i%400 == 0) {
			daycnt += 366
		} else {
			daycnt += 365
		}
	}
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		days[1] = 29
	}
	for i := 0; i < month-1; i++ {
		daycnt += days[i]
	}
	daycnt += day
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return names[(daycnt+4)%7]
}