# 17. 电话号码的字母组合

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/) |
| 难度 | Medium |
| 分类 | backtracking |
| 标签 | backtracking, string |
| 文件 | `backtracking/code_17.go` |

## 题目描述

给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。数字到字母的映射与电话按键相同（与电话按键相同）。注意 `1` 不对应任何字母。

**示例 1：**

```
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**示例 2：**

```
输入：digits = ""
输出：[]
```

**示例 3：**

```
输入：digits = "2"
输出：["a","b","c"]
```

## 解题算法

### 方法：回溯

**核心思路：**
使用回溯法生成所有可能的字母组合。首先建立数字到字母的映射表，然后递归地为每个数字选择对应的字母。在每一层递归中，处理当前数字对应的所有字母，选择一个字母加入路径，然后递归处理下一个数字。当路径长度等于输入字符串长度时，将当前组合加入结果集。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(4^n * n)，其中 n 是输入字符串的长度，每个数字最多对应 4 个字母
- 空间复杂度：O(n)，递归栈的深度为 n
