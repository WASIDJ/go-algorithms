# 提交信息规范

## 格式

```
<type>(<scope>): <subject>
// 空一行
<body> // 可选
```

所有部分使用**中文**描述（subject 和 body 均用中文）。

## 提交类型

| 类型 | 适用场景 | 示例 |
|------|----------|------|
| `algo` | 新增或优化算法题解 | `algo(two_sum): 哈希表解法 O(n)` |
| `note` | 读书笔记 / 知识点整理 | `note(408): 操作系统-进程调度总结` |
| `fix` | 修复代码错误 | `fix(linked_list): 处理空指针边界 case` |
| `refactor` | 重构代码，不改变逻辑 | `refactor: 抽取公共 ListNode 工具函数` |
| `test` | 增删改测试 | `test: two_sum 补充多个用例` |
| `docs` | 文档/README 相关 | `docs: README 更新刷题进度` |
| `chore` | 构建/配置/工具链变动 | `chore: 移除 submodule 改为本地目录跟踪` |
| `perf` | 性能优化（降复杂度等） | `perf(group_anagrams): 计数数组替代哈希表` |
| `project` | 项目代码（非算法题） | `project: 初始化项目脚手架` |
| `study` | 408 / 计算机基础复习 | `study(network): TCP 拥塞控制总结` |
| `cert` | 考证相关材料 | `cert: AWS 备考笔记更新` |

## Scope 建议

- 题解用题目英文名（snake_case）：`algo(two_sum)`、`algo(group_anagrams)`
- 笔记用科目缩写：`note(os)`、`note(计网)`、`note(408)`
- 基础复习用：`study(os)`、`study(网络)`、`study(数据库)`
- 修改公共模块用模块名：`refactor(list_node)`、`refactor(tree_node)`

## 规则

- subject 首字母小写，末尾不加句号，不超过 **72 字符**
- body 说明**为什么改**而不是改了什么（diff 已经说明改了什么）
- 同一道题的多次迭代用相同 scope 聚合，方便 `git log --grep` 检索
- 如果需要升级复杂度（如 O(n²)→O(n)），在 body 说明 trade-off

## 示例

### 新增题解

```
algo(two_sum): 哈希表解法 O(n)

用 map[值]下标 记录已遍历元素，一遍遍历即可找到两数之和。
优于暴力 O(n²) 解法。
```

### 优化题解

```
perf(group_anagrams): 用 [26]int 计数代替排序

对每个字符串用固定长度数组计数，避免 sort.Slice 的 O(k log k) 开销。
整体从 O(n k log k) 降为 O(n k)。
```

### 笔记

```
note(os): 进程与线程对比总结

整理 PCB、上下文切换开销、共享资源的对比表，方便面试前翻看。
```

### 多项改动

```
algo: 完成 Hot 100 前 20 题

包含：two_sum、group_anagrams、merge_two_lists、preorder_traversal
```

### 修复

```
fix(merge_two_lists): 处理 l1/l2 为空指针 case

LeetCode 测试通过但在本地运行 panic，原因是未判断 nil。
```