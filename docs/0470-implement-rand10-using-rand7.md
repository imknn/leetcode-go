# 470. 用 Rand7() 实现 Rand10()

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [470. 用 Rand7() 实现 Rand10()](https://leetcode.cn/problems/implement-rand10-using-rand7/) |
| 难度 | Medium |
| 分类 | random |
| 标签 | random, rejection sampling |
| 文件 | `random/code_470.go` |

## 题目描述

已有方法 `rand7()` 可生成 1 到 7 范围内的随机整数，写出一个方法 `rand10()` 生成 1 到 10 范围内的随机整数。

不要使用系统的 `Math.random()` 方法。

**示例 1:**

```
输入: 1
输出: [7]
```

**示例 2:**

```
输入: 2
输出: [8,4]
```

**示例 3:**

```
输入: 3
输出: [8,1,10]
```

## 解题算法

### 方法：拒绝采样

**核心思路：**
利用两次 `rand7()` 调用构造一个均匀分布的随机数。`(rand7()-1)*7 + rand7() - 1` 可以生成 `[0, 48]` 范围内的均匀随机数。由于 49 不是 10 的倍数，直接取模会引入偏差，因此只接受 `[0, 39]` 范围内的值（共 40 个），拒绝 `[40, 48]` 范围的值。对接受的值取模 10 再加 1，即可得到 `[1, 10]` 的均匀随机数。

**代码实现：**
```go
func rand7() int {
	time.Sleep(time.Duration(1) * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}

func rand10() int {
	var m int
	for {
		m = (rand7()-1)*7 + rand7() - 1
		if m <= 39 {
			break
		}
	}
	return m%10 + 1
}
```

**复杂度分析：**
- 时间复杂度：期望 O(1)，每次循环有 40/49 的概率被接受，平均调用约 2.45 次 `rand7()`
- 空间复杂度：O(1)，只使用了常数级别的额外空间
