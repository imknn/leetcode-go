package backtracking

/* 78. 子集 */
func subsets(nums []int) [][]int {
	// 幂
	length := exp2(len(nums))
	var result = make([][]int, 0)
	for i := 0; i < length; i++ {
		result = append(result, make([]int, 0))
		for j, num := range nums {
			// 使用位运算标记当前值是否包含在里面
			if 1<<uint32(j)&i != 0 {
				result[i] = append(result[i], num)
			}
		}
	}
	return result
}

func exp2(n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num *= 2
	}
	return num
}
