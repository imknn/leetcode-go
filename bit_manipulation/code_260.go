package bit_manipulation

/* 260. 只出现一次的数字 III */
func singleNumber03(nums []int) []int {
	result := make([]int, 2)
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}
	i := 0
	for k, v := range m {
		if v == 1 {
			result[i] = k
			i++
		}
		if i == 2 {
			break
		}
	}
	return result
}
