---
name: solve-leetcode
description: 解决 LeetCode 题目 - 自动获取题目、编写代码、生成文档
user_invocable: true
---

# LeetCode 题目解决技能

当用户提供 LeetCode 题目链接时，自动执行以下流程：

## 步骤

### 1. 获取题目信息
- 用 Playwright MCP 打开链接
- 用 `browser_evaluate` 提取题目描述：
  ```js
  document.querySelector('[data-track-load="description_content"]')?.innerText
  ```
- 提取题号、题目名称、难度、标签、约束条件
- 关闭浏览器

### 2. 确定分类目录
根据题目标签和内容，将题目归入对应分类：
- `string/` — 字符串相关
- `array/` — 数组相关
- `linked_list/` — 链表相关
- `stack_queue/` — 栈/队列相关
- `binary_search/` — 二分查找
- `dynamic_programming/` — 动态规划
- `bit_manipulation/` — 位运算
- `math/` — 数学
- `hash/` — 哈希表
- `backtracking/` — 回溯
- `sorting/` — 排序
- `random/` — 随机化
- `other/` — 其他

### 3. 编写代码
- 在对应目录创建 `code_<题号>.go`
- 包名与目录名一致（如 `package string`）
- 编写最优解法

### 4. 编写测试
- 创建 `code_<题号>_test.go`
- 使用表驱动测试格式
- 覆盖示例用例和边界情况
- 运行测试确保通过

### 5. 更新 README
- 更新对应分类目录的 `README.md`
- 将题目状态标记为"已解决"

### 6. 生成文档
- 在 `docs/` 目录创建 `题号-题目英文名.md`
- 文档包含：
  - 题目信息表格（含 LeetCode 链接）
  - 题目描述和示例
  - 解题算法详解
  - 代码实现
  - 复杂度分析

## 文档模板

```markdown
# 题号. 题目名称

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [题号. 题目名称](https://leetcode.cn/problems/题目英文名/) |
| 难度 | 简单/中等/困难 |
| 分类 | 目录名 |
| 标签 | 标签1、标签2 |
| 文件 | `目录/code_题号.go` |

## 题目描述

（题目内容）

## 解题算法

### 方法：算法名称

**核心思路：**
（思路说明）

**代码实现：**
（代码）

**复杂度分析：**
- 时间复杂度：O(?)
- 空间复杂度：O(?)
```
