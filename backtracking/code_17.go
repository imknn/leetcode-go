package backtracking

// 17. 电话号码的字母组合
// 回溯法
// 时间 O(4^n · n)，空间 O(n)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	keyMap := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var result []string

	var backtrack func(index int, path []byte)
	backtrack = func(index int, path []byte) {
		if index == len(digits) {
			combo := make([]byte, len(path))
			copy(combo, path)
			result = append(result, string(combo))
			return
		}
		letters := keyMap[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []byte{})
	return result
}
