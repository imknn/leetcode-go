package array

/* 27. 移除元素*/
func removeElement(nums []int, val int) int {
	length := len(nums)
	tail := length - 1
	for i := 0; i <= tail; {
		if nums[i] == val {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			i++
		}
	}
	return tail + 1
}
