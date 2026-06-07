# 11. 盛最多水的容器

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/) |
| 难度 | Medium |
| 分类 | array |
| 标签 | two pointers |
| 文件 | `array/code_11.go` |

## 题目描述

给定一个长度为 `n` 的整数数组 `height`。其中 `height[i]` 表示第 `i` 条线的高度。

从其中选择两条线，与 x 轴构成一个容器，容器可以容纳水。

返回容器可以存储的最大水量。

**示例 1：**

```
输入：height = [1,8,6,2,5,4,8,3,7]
输出：49
解释：选择第 2 条线和第 9 条线，宽度为 8，高度为 min(8,7) = 7，面积为 8 * 7 = 56。
```

**示例 2：**

```
输入：height = [1,1]
输出：1
```

## 解题算法

### 方法：双指针

**核心思路：**
使用左右两个指针分别指向数组的两端。每次计算当前两个指针所构成的容器面积，然后移动较短一侧的指针。因为容器的面积由较短的线决定，移动较长一侧只会让宽度变小而高度不变或变小，不可能得到更大的面积。移动较短一侧则有可能遇到更高的线，从而可能获得更大的面积。

**代码实现：**
```go
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)
