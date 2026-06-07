# 728. 自除数

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [728. 自除数](https://leetcode.cn/problems/self-dividing-numbers/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学 |
| 文件 | `math/code_728.go` |

## 题目描述

自除数是指可以被它包含的每一位数除尽的数。

例如，128 是一个自除数，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。

自除数不允许包含 0。给定 left 和 right 数字列表，计算范围内（包含 left 和 right）自除数的个数。

**示例 1：**
```
输入：left = 1, right = 22
输出：9
解释：自除数有 [1, 2, 3, 4, 5, 6, 7, 8, 9]。
```

## 解题算法

### 方法：逐个检查

**核心思路：**
遍历 [left, right] 范围内的每个数，逐位提取数字进行检查。如果某一位是 0 则直接跳过（不允许包含 0），如果某一位不能整除原数也跳过。如果所有位都通过检查，则该数是自除数。

**代码实现：**
```go
func selfDividingNumbers(left int, right int) []int {
	m := make([]int, 0)
	for i := left; i <= right; i++ {
		a := i
		for a > 0 {
			b := a % 10
			if b == 0 {
				break
			} else if b != 1 && i%b != 0 {
				break
			}
			a /= 10
		}
		if a == 0 {
			m = append(m, i)
		}
	}
	return m
}
```

**复杂度分析：**
- 时间复杂度：O(n * d)，n 为范围内的数字个数，d 为每个数的位数
- 空间复杂度：O(1)（不计输出数组）
