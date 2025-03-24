# 修改上次提交的 committer date

## 一、命令介绍

该命令用于修改最近一次提交的 committer date（提交者日期）。通过设置环境变量 `GIT_COMMITTER_DATE` 并重新提交，可以更改最后一次提交的时间记录。

## 二、命令使用

### 命令格式

```bash
GIT_COMMITTER_DATE="date" git commit --amend
```

### 命令示例

```bash
GIT_COMMITTER_DATE="2023-10-01T12:00:00" git commit --amend
```

## 三、输出说明

执行该命令后，Git 会重新提交最后一次提交，并使用指定的 `GIT_COMMITTER_DATE` 作为新的提交者日期。不会产生额外的输出信息，但可以使用 `git log` 查看提交记录以确认日期是否已更新。

## 四、注意事项

### 日期格式

`GIT_COMMITTER_DATE` 的日期格式应符合 ISO 8601 标准，例如 `YYYY-MM-DDTHH:MM:SS`。

### 提交记录

使用该命令会修改提交记录中的 committer date，但不会更改提交的内容或作者信息。请确保在执行前了解其影响。

### 仓库状态

执行该命令会修改提交历史，因此不推荐在已经推送到远程仓库的提交上使用。如果需要修改远程仓库的提交记录，请谨慎操作并通知团队成员。