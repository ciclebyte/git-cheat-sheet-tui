# 列出 repository 配置

## 一、命令介绍

`git config --local --list` 命令用于列出当前 Git 仓库的本地配置信息。本地配置信息存储在仓库的 `.git/config` 文件中，仅对当前仓库有效。

## 二、命令使用

### 命令格式

```bash
git config --local --list
```

### 命令示例

```bash
git config --local --list
```

## 三、输出说明

以下是执行该命令后可能输出的配置信息及对应解释：

```bash
# 仓库的格式版本，0 表示默认的格式
core.repositoryformatversion=0
# 是否跟踪文件模式，在 Windows 系统中通常设置为 false
core.filemode=false
# 是否为裸仓库，裸仓库没有工作目录，通常用于服务器端
core.bare=false
# 是否记录所有引用的更新，设置为 true 可以方便查看引用的变化
core.logallrefupdates=true
# 是否支持符号链接，设置为 false 表示不支持
core.symlinks=false
# 是否忽略文件名大小写，设置为 true 表示忽略
core.ignorecase=true
# 远程仓库 origin 的 URL 地址
remote.origin.url=http://192.168.31.93:3003/xxx/xxx.git
# 从远程仓库 origin 拉取的引用范围
remote.origin.fetch=+refs/heads/*:refs/remotes/origin/*
# master 分支对应的远程仓库为 origin
branch.master.remote=origin
# master 分支合并时对应的远程分支
branch.master.merge=refs/heads/master
# VSCode 合并时的基础分支为 origin/master
branch.master.vscode-merge-base=origin/master
```

## 四、注意事项

### 仓库必须已初始化

该命令仅在已经初始化的 Git 仓库中有效。如果当前目录不是一个 Git 仓库，执行该命令将不会输出任何信息。

### 本地配置优先级

Git 配置分为系统级、全局级和本地级，本地配置优先级最高。如果在不同级别配置了相同的变量，本地配置将覆盖全局和系统配置。

### 网络安全

如果远程仓库的 URL 是通过 HTTP 协议访问的，建议使用 HTTPS 协议以提高安全性。
