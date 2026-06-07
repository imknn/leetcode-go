# 665. 非递减数列

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [665. 非递减数列](https://leetcode.cn/problems/non-decreasing-array/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | greedy |
| 文件 | `array/code_665.go` |

## 题目描述

给你一个长度为 `n` 的整数数组，请你判断在 **最多** 改变 `1` 个元素的情况下，该数组能否变成一个非递减数列。

非递减数列定义为：数组中所有 `i` 的总和（`0 <= i <= n-2`）都满足 `nums[i] <= nums[i+1]`。

**示例 1：**

```
输入：nums = [4,2,3]
输出：true
解释：你可以把第一个 4 变成 1 从而得到一个非递减数列。
```

**示例 2：**

```
输入：nums = [4,2,1]
输出：false
```

## 解题算法

### 方法：贪心（逐次修正）

**核心思路：**
遍历数组，当发现 `nums[j-1] > nums[j]` 时，需要进行修正。有两种选择：
1. 将 `nums[j-1]` 降低为 `nums[j]`（小值）
2. 将 `nums[j]` 提升为 `nums[j-1]`（大值）

先尝试第一种方案，修改后检查整个数组是否非递减。如果不行，则尝试第二种方案。由于只允许修改一次，遇到第一次逆序后就需要做出正确选择。

**代码实现：**
```go
func checkPossibility(nums []int) bool {
	length := len(nums)
	for j := 1; j < length; j++ {
		if nums[j-1] > nums[j] {
			l := nums[j-1]
			r := nums[j]
			// 统一设置成右值（小值）
			nums[j-1] = r
			nums[j] = r
			if !isAsc(nums, length) {
				// 统一设置成左值 （大值）
				nums[j-1] = l
				nums[j] = l
				return isAsc(nums, length)
			}
		}
	}
	return true
}

// 判断数列是否递增
func isAsc(nums []int, length int) bool {
	for i := 1; i < length; i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)
