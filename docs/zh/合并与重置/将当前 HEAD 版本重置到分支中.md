# 将当前 HEAD 版本重置到分支中

## 一、命令介绍

`git rebase` 命令用于将当前分支的提交记录重新应用到指定的分支上。它可以用来整理提交历史，将分支的提交记录转移到另一个分支的顶部。

## 二、命令使用

### 命令格式

```bash
git rebase <branch>
```

### 命令示例

```bash
git rebase main
```

## 三、注意事项

### 请勿重置已发布的提交

在多人协作的仓库中，避免对已经推送到远程仓库的提交进行 `rebase` 操作，因为这会导致历史记录的变更，可能会对其他开发者的工作造成影响。

### 合并冲突

在执行 `rebase` 过程中，如果遇到冲突，Git 会暂停 `rebase` 过程并提示你解决冲突。解决冲突后，使用 `git rebase --continue` 继续 `rebase` 过程。

### 中止 `rebase`

如果你在 `rebase` 过程中遇到问题，可以使用以下命令中止 `rebase` 操作：

```bash
git rebase --abort
```

### 交互式 `rebase`

如果你想在 `rebase` 过程中对提交进行编辑、合并或删除等操作，可以使用交互式 `rebase`：

```bash
git rebase -i <branch>
```

### 强制推送

在 `rebase` 完成后，如果要将更改推送到远程仓库，可能需要使用 `--force` 或 `--force-with-lease` 选项，因为 `rebase` 会改变提交历史：

```bash
git push --force-with-lease
```