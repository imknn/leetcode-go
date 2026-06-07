# 118. 杨辉三角

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [118. 杨辉三角](https://leetcode.cn/problems/pascals-triangle/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | DP |
| 文件 | `array/code_118.go` |

## 题目描述

给定一个非负整数 `numRows`，生成杨辉三角的前 `numRows` 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

**示例 1：**

```
输入：numRows = 5
输出：
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

**示例 2：**

```
输入：numRows = 1
输出：[[1]]
```

## 解题算法

### 方法：动态规划（利用对称性）

**核心思路：**
逐行生成杨辉三角。利用杨辉三角的对称性，每次只计算左半部分的值，然后通过对称复制到右半部分。每一行的第一个和最后一个元素都是 1，中间的元素由上一行相邻两个元素相加得到。

对于第 `r` 行（从 3 开始），计算的个数为 `r/2`（偶数行）或 `(r+1)/2`（奇数行），利用对称性 `row[r-1-c] = row[c]` 完成右半部分的填充。

**代码实现：**
```go
func generate(numRows int) [][]int {
	result := make([][]int, 0)
	if numRows == 0 {
		return result
	}
	result = append(result, []int{1})
	if numRows == 1 {
		return result
	}
	result = append(result, []int{1, 1})
	if numRows == 2 {
		return result
	}
	for r := 3; r <= numRows; r++ {
		row := make([]int, r)
		count := 0
		if r&1 == 1 {
			count = (r + 1) / 2
		} else {
			count = r / 2
		}
		for c := 0; c < count; c++ {
			if c == 0 {
				row[c] = 1
				row[r-1-c] = row[c]
			} else {
				row[c] = result[r-2][c-1] + result[r-2][c]
				row[r-1-c] = row[c]
			}
		}
		result = append(result, row)
	}
	return result
}
```

**复杂度分析：**
- 时间复杂度：O(numRows^2)
- 空间复杂度：O(1)（不考虑输出数组的空间）
