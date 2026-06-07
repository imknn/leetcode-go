package array

func getRow(rowIndex int) []int {
	rowIndex++
	if rowIndex == 1 {
		return []int{1}
	}
	if rowIndex == 2 {
		return []int{1, 1}
	}
	row := make([]int, rowIndex)
	row[0] = 1
	row[1] = 1
	for r := 3; r <= rowIndex; r++ {
		count := 0
		if r&1 == 1 {
			count = (r + 1) / 2
		} else {
			count = r / 2
		}
		t := 0
		for c := 0; c < count; c++ {
			t, row[c] = row[c], t+row[c]
			row[r-1-c] = row[c]
		}
	}
	return row
}
