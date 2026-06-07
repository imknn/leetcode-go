package backtracking

// 784. 字母大小写全排列
// 对每个字母尝试大写和小写两种情况，数字直接跳过
// 使用回溯/DFS 生成所有排列

func letterCasePermutation(s string) []string {
	var result []string

	var dfs func(index int, path []byte)
	dfs = func(index int, path []byte) {
		if index == len(s) {
			result = append(result, string(path))
			return
		}

		// 数字直接加入，跳过
		if s[index] >= '0' && s[index] <= '9' {
			dfs(index+1, append(path, s[index]))
			return
		}

		// 字母：尝试小写
		dfs(index+1, append(path, s[index]|32))
		// 字母：尝试大写
		dfs(index+1, append(path, s[index]&^32))
	}

	dfs(0, []byte{})
	return result
}
