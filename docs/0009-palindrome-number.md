# 9. 回文数

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [9. 回文数](https://leetcode.cn/problems/palindrome-number/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学 |
| 文件 | `math/code_9.go` |

## 题目描述

给你一个整数 x，如果 x 是一个回文整数，返回 true；否则，返回 false。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。比如 121 是回文数，而 123 不是。

**示例 1：**
```
输入：x = 121
输出：true
```

**示例 2：**
```
输入：x = -121
输出：false
解释：从左向右读，为 -121 。 从右向左读，为 121- 。因此它不是一个回文数。
```

**示例 3：**
```
输入：x = 10
输出：false
解释：从右向左读，为 01 。因此它不是一个回文数。
```

## 解题算法

### 方法：逐位提取比较

**核心思路：**
将整数的每一位数字依次提取并存入数组，然后通过双指针从两端向中间比较每一位是否相同。负数直接返回 false。

**代码实现：**
```go
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var num []int
	for x > 0 {
		num = append(num, x%10)
		x /= 10
	}

	for l, i := len(num), 0; i < l/2; i++ {
		if num[i] != num[l-1-i] {
			return false
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(log(x))，需要遍历所有位数
- 空间复杂度：O(log(x))，需要数组存储每一位数字
