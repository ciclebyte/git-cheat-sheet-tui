# 从远程仓库下载所有修改，但不合并到 HEAD 中

## 一、命令介绍

`git fetch` 命令用于从远程仓库下载所有修改，但不会自动将这些修改合并到当前分支的 HEAD 中。这个命令通常用于查看远程仓库的更新情况，而不会直接影响本地的工作分支。

## 二、命令使用

### 命令格式

```bash
git fetch <remote>
```

### 命令示例

```bash
git fetch origin
```

## 三、输出说明

以下是执行 `git fetch` 命令后可能输出的信息及解释：

```bash
From https://github.com/user/repo
 * [new branch]      feature-branch -> origin/feature-branch
```

- `From https://github.com/user/repo` 表示远程仓库的 URL。
- `* [new branch] feature-branch -> origin/feature-branch` 表示远程仓库中有新的分支 `feature-branch`，并且它被下载到本地的 `origin/feature-branch` 引用中。

## 四、注意事项

### 远程仓库名称

在执行 `git fetch` 命令时，需要指定远程仓库的名称（如 `origin`）。如果未指定远程仓库名称，默认会从所有远程仓库下载修改。

### 不自动合并

`git fetch` 命令不会自动将下载的修改合并到当前分支的 HEAD 中。如果需要将这些修改合并到当前分支，可以使用 `git merge` 或 `git rebase` 命令。