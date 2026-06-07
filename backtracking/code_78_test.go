package backtracking

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{9, 0, 3, 5, 7}
	result := subsets(nums)
	fmt.Println(result)
	if len(result) != 32 {
		t.Fail()
	}
}
