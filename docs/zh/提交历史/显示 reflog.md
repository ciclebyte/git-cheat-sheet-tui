# 显示 reflog

## 一、命令介绍

`git reflog` 命令用于显示本地仓库的引用日志（reflog）。引用日志记录了 HEAD 和分支的更新历史，包括提交、重置、合并等操作。即使某些提交不再被任何分支引用，仍然可以通过 reflog 查看这些提交的历史。

## 二、命令使用

### 命令格式

```bash
git reflog show [<ref>]
```

### 命令示例

```bash
git reflog show
```

### 输出说明

以下是执行该命令后可能输出的日志信息及对应解释：

```bash
# 示例输出
abcd123 HEAD@{0}: commit: Update README file
efgh456 HEAD@{1}: reset: moving to HEAD~1
ijkl789 HEAD@{2}: commit: Add new feature
```

- `abcd123`: 提交的哈希值。
- `HEAD@{0}`: 引用的最新状态。
- `commit: Update README file`: 操作类型及描述。

## 四、注意事项

### reflog 的有效期

`git reflog` 记录的日志默认保留 90 天。过期后，日志条目将被自动删除。

### 适用于本地仓库

`git reflog` 仅适用于本地仓库，不会影响远程仓库的引用日志。