# 896. 单调数列

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [896. 单调数列](https://leetcode.cn/problems/monotonic-array/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | traversal |
| 文件 | `array/code_896.go` |

## 题目描述

如果数组是单调递增或单调递减的，那么它是 **单调数组** 的。

形式化定义：如果对于所有的 `i`，都有 `nums[i] <= nums[i+1]`，那么数组 `nums` 是单调递增的。如果对于所有的 `i`，都有 `nums[i] >= nums[i+1]`，那么数组 `nums` 是单调递减的。

**示例 1：**

```
输入：nums = [1,2,2,3]
输出：true
```

**示例 2：**

```
输入：nums = [6,5,4,4]
输出：true
```

**示例 3：**

```
输入：nums = [1,3,2]
输出：false
```

## 解题算法

### 方法：一次遍历记录状态

**核心思路：**
使用一个 `status` 变量记录当前的单调状态：0 表示相等，1 表示递增，2 表示递减。遍历数组时：
- 如果 `A[i-1] == A[i]`，跳过（不影响单调性）
- 如果 `A[i-1] < A[i]` 且当前不是递减状态，则标记为递增
- 如果 `A[i-1] > A[i]` 且当前不是递增状态，则标记为递减
- 如果与当前状态矛盾，则不是单调数列

**代码实现：**
```go
func isMonotonic(A []int) bool {
	status, length := 0, len(A) // status: 0相等，1递增，2递减
	if length <= 1 {
		return true
	}
	for i := 1; i < length; i++ {
		if A[i-1] == A[i] {
			continue
		}
		if A[i-1] < A[i] && status != 2 {
			status = 1
			continue
		}
		if A[i-1] > A[i] && status != 1 {
			status = 2
			continue
		}
		return false
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)
