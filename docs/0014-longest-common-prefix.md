# 14. 最长公共前缀

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [14. 最长公共前缀](https://leetcode.cn/problems/longest-common-prefix/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_14.go` |

## 题目描述

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 `""`。

**示例 1：**
```
输入：strs = ["flower","flow","flight"]
输出："fl"
```

**示例 2：**
```
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
```

**提示：**
- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` 仅由小写英文字母组成

## 解题算法

### 方法：纵向扫描

**核心思路：**
以第一个字符串为基准，逐个字符与其余字符串比较。对于每个位置，如果所有字符串在该位置的字符都相同，则该字符属于公共前缀；一旦发现不匹配或某个字符串已结束，就截断到目前为止的公共前缀。

**代码实现：**
```go
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	seekStr := strs[0]
	for i := 1; i < length; i++ {
		length1 := len(seekStr)
		length2 := len(strs[i])
		if length1 > length2 {
			seekStr, strs[i] = strs[i], seekStr
			length1, length2 = length2, length1
		}
		for j := 0; j < length1; j++ {
			if seekStr[j] != strs[i][j] {
				if j == 0 { // 第一个就不等，直接返回
					return ""
				}
				seekStr = seekStr[0:j]
				break
			}
		}
	}
	return seekStr
}
```

**复杂度分析：**
- 时间复杂度：O(S)，其中 S 是所有字符串字符数之和
- 空间复杂度：O(1)，只使用常数个变量（不计结果字符串）
