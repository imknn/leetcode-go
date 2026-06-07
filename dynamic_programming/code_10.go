package dynamic_programming

// 10. 正则表达式匹配
// dp[i][j] 表示 s[0..i-1] 与 p[0..j-1] 是否匹配
// 时间 O(mn)，空间 O(mn)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[i][j]: s 前 i 个字符是否匹配 p 前 j 个字符
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 空 s 匹配 p 中 '*' 使前面对应字符被消掉的情况
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// '*' 匹配零次：dp[i][j-2]
				dp[i][j] = dp[i][j-2]
				// '*' 匹配一次或多次：s[i-1] 必须匹配 p[j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
