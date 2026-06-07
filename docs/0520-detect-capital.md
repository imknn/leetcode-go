# 520. 检测大写字母

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [520. 检测大写字母](https://leetcode.cn/problems/detect-capital/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_520.go` |

## 题目描述

给定一个单词 `word`，如果大写字母使用得当则认为其满足以下条件：

1. 所有字母都是大写的，例如 `"USA"`
2. 所有字母都是小写的，例如 `"leetcode"`
3. 仅首字母大写，例如 `"Google"`

否则，认为其使用不当，例如 `"FlaG"`。

给定一个字符串 `word`，返回如果大写字母使用得当则返回 `true`，否则返回 `false`。

**示例 1：**
```
输入：word = "USA"
输出：true
```

**示例 2：**
```
输入：word = "FlaG"
输出：false
```

**提示：**
- `1 <= word.length <= 100`
- `word` 由小写和大写英文字母组成

## 解题算法

### 方法：统计大写字母个数

**核心思路：**
统计大写字母的个数。满足以下任一条件即为合法：
1. 大写字母数为 0（全小写）
2. 大写字母数等于字符串长度（全大写）
3. 大写字母数为 1 且首字母为大写（仅首字母大写）

**代码实现：**
```go
func detectCapitalUse(word string) bool {
	count := 0
	// 统计大写字母
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}
	if count == 0 || count == len(word) || (count == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历字符串统计大写字母
- 空间复杂度：O(1)，只使用了常数个变量
