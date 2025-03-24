# 以 rebase 方式将远程分支与本地合并

## 一、命令介绍

`git pull --rebase` 命令用于将远程分支的更改以 rebase 的方式合并到本地分支。与普通的 `git pull` 不同，`git pull --rebase` 会将本地的提交“基变”到远程分支的最新提交之上，从而保持提交历史的线性。

## 二、命令使用

### 命令格式

```bash
git pull --rebase <remote> <branch>
```

### 命令示例

```bash
git pull --rebase origin main
```

## 三、输出说明

执行该命令后，Git 会将远程分支 `main` 的更改拉取到本地，并将本地的提交依次应用到这些更改之上。如果存在冲突，Git 会暂停 rebase 过程并提示你解决冲突。解决冲突后，可以使用 `git rebase --continue` 继续 rebase 过程。

## 四、注意事项

### 冲突处理

在执行 `git pull --rebase` 时，如果遇到冲突，Git 会暂停 rebase 过程并提示你解决冲突。解决冲突后，需要使用 `git add` 将更改标记为已解决，然后使用 `git rebase --continue` 继续 rebase 过程。如果想要取消 rebase，可以使用 `git rebase --abort`。

### 提交历史

使用 `git pull --rebase` 可以保持提交历史的线性，但也会改变本地提交的 SHA-1 值。因此，如果你已经将本地分支推送到远程仓库，使用 `git pull --rebase` 可能会导致冲突或其他问题。在这种情况下，建议谨慎操作或与团队成员协商后再进行 rebase。