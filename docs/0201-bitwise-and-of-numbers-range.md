# 201. 数字范围按位与

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [201. 数字范围按位与](https://leetcode.cn/problems/bitwise-and-of-numbers-range/) |
| 难度 | 中等 |
| 分类 | bit_manipulation |
| 标签 | 位运算 |
| 文件 | `bit_manipulation/code_201.go` |

## 题目描述

给你两个整数 `left` 和 `right`，表示区间 `[left, right]`，返回此区间内所有数字按位与的结果（包含端点 `left` 和 `right`）。

**示例 1：**
```
输入：left = 5, right = 7
输出：4
```

**示例 2：**
```
输入：left = 0, right = 0
输出：0
```

**示例 3：**
```
输入：left = 1, right = 2147483647
输出：0
```

**提示：**
- `0 <= left <= right <= 2^31 - 1`

## 解题算法

### 方法：逐个按位与

**核心思路：**
从 `left` 开始，依次与 `[left+1, right]` 范围内的每个数字进行按位与运算。当遇到 `left == 0` 时直接返回 0，因为 0 与任何数按位与结果都是 0。

**代码实现：**
```go
func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	num := m
	for i := m + 1; i <= n; i++ {
		num &= i
	}
	return num
}
```

**复杂度分析：**
- 时间复杂度：O(n - m)，遍历区间内所有数字
- 空间复杂度：O(1)，只使用常量额外空间
