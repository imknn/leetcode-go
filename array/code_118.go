package array

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}
	for r := 3; r <= numRows; r++ {
		row := make([]int, r)
		count := 0
		if r&1 == 1 {
			count = (r + 1) / 2
		} else {
			count = r / 2
		}
		for c := 0; c < count; c++ {
			if c == 0 {
				row[c] = 1
				row[r-1-c] = row[c]
			} else {
				row[c] = result[r-2][c-1] + result[r-2][c]
				row[r-1-c] = row[c]
			}
		}
		result = append(result, row)
	}
	return result
}
