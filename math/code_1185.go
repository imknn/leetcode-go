package math

/* 1185. 一周中的第几天  */
func dayOfTheWeek(day int, month int, year int) string {
	week := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if month <= 2 {
		month += 12
		year--
	}
	t := (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1) % 7
	return week[t]
}
