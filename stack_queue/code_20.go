package stack_queue

/* 20. 有效的括号 */
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	if length%2 != 0 {
		return false
	}
	// 栈
	top := -1 // 栈顶索引
	save := make([]byte, length)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(', '[', '{':
			// 入栈
			top++
			save[top] = s[i]
		case ')':
			if top < 0 || save[top] != '(' {
				return false
			} else {
				top--
			}
		case ']':
			if top < 0 || save[top] != '[' {
				return false
			} else {
				top--
			}
		case '}':
			if top < 0 || save[top] != '{' {
				return false
			} else {
				top--
			}
		}
	}
	return top == -1
}
