package array

// 896. 单调数列
func isMonotonic(A []int) bool {
	status, length := 0, len(A) // status: 0相等，1递增，2递减
	if length <= 1 {
		return true
	}
	for i := 1; i < length; i++ {
		if A[i-1] == A[i] {
			continue
		}
		if A[i-1] < A[i] && status != 2 {
			status = 1
			continue
		}
		if A[i-1] > A[i] && status != 1 {
			status = 2
			continue
		}
		return false
	}
	return true
}
