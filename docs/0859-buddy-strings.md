# 859. 亲密字符串

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [859. 亲密字符串](https://leetcode.cn/problems/buddy-strings/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 哈希表、字符串 |
| 文件 | `string/code_859.go` |

## 题目描述

给你两个字符串 `s` 和 `goal`，如果可以通过交换 `s` 中的两个字母得到与 `goal` 相等的结果，那么返回 `true`。否则返回 `false`。

交换定义为：选中两个下标 `i` 和 `j`（下标从 0 开始），然后交换 `s[i]` 和 `s[j]`。每个下标最多只能交换一次。

**示例 1：**
```
输入：s = "ab", goal = "ba"
输出：true
解释：你可以交换 s[0]='a' 和 s[1]='b' 得到 "ba"。
```

**示例 2：**
```
输入：s = "ab", goal = "ab"
输出：false
解释：你只能交换 s[0]='a' 和 s[1]='b' 一次，结果还是 "ab"。
```

**示例 3：**
```
输入：s = "aa", goal = "aa"
输出：true
解释：你可以交换 s[0]='a' 和 s[0]='a' 两次，得到 "aa"。
```

**提示：**
- `1 <= s.length, goal.length <= 2 * 10^4`
- `s` 和 `goal` 由小写英文字母组成

## 解题算法

### 方法：位运算 + 差异统计

**核心思路：**
收集 s 和 goal 中字符不同的位置。如果：
- 差异位置为 0：要求 s 中有重复字符（可以自己交换自己），用位运算检测是否有重复字母
- 差异位置恰好为 2：检查交换后是否相等，即 s[diff[0]] == goal[diff[1]] 且 s[diff[1]] == goal[diff[0]]
- 其他情况均返回 false

**代码实现：**
```go
func buddyStrings(A string, B string) bool {
	diff := make([]int, 0)
	m := len(A)
	if m != len(B) {
		return false
	}
	check, hasSame := 0, false
	for i := 0; i < m; i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
		if !hasSame && ((check>>(A[i]-'a'))&1) == 1 {
			hasSame = true
		}
		if !hasSame {
			check = check | (1 << (A[i] - 'a'))
		}
	}
	n := len(diff)
	if n == 0 {
		return hasSame
	} else if n == 2 {
		return A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]]
	}
	return false
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历一次字符串
- 空间复杂度：O(n)，存储差异位置列表（最坏情况）
