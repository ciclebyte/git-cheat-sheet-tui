# 将分支合并到当前 HEAD 中

## 一、命令介绍

`git merge` 命令用于将指定分支的更改合并到当前所在的分支（即 HEAD 指向的分支）中。该命令常用于在开发过程中，将一个分支的功能或修复合并到主分支或其他目标分支。

## 二、命令使用

### 命令格式

```bash
git merge <branch>
```

### 命令示例

```bash
git merge feature-branch
```

## 三、输出说明

以下是执行该命令后可能输出的信息及对应解释：

```bash
# 如果合并成功
Merge made by the 'recursive' strategy.
 file1.txt | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

# 如果发生冲突
CONFLICT (content): Merge conflict in file1.txt
Automatic merge failed; fix conflicts and then commit the result.
```

## 四、注意事项

### 合并冲突

如果在合并过程中发生冲突，Git 会中断合并操作并在冲突文件中标记出冲突的部分。此时需要手动解决冲突，并使用 `git add` 命令将解决后的文件标记为已解决，然后再使用 `git commit` 完成合并。

### 合并策略

`git merge` 默认使用递归策略进行合并。可以通过 `-s` 选项指定其他合并策略，例如 `-s octopus`（用于合并多个分支）或 `-s ours`（保留当前分支的更改）。