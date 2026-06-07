package stack_queue

import (
	"bytes"
)

/* 1021. 删除最外层的括号 */
func removeOuterParentheses(S string) string {
	var buffer bytes.Buffer
	left, size := 0, 0
	for idx, ch := range S {
		if ch == '(' {
			size++
		} else {
			size--
		}
		if size == 0 {
			// 原语 S[left : idx+1]), 去除前后括号就是S[left+1 : idx]
			buffer.WriteString(S[left+1 : idx])
			left = idx + 1
		}
	}
	return buffer.String()
}
