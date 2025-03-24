# 设置 git 命令行输出为彩色

## 一、命令介绍

`git config --global color.ui auto` 命令用于设置 Git 命令行输出的颜色为自动模式。启用此功能后，Git 会根据不同的操作和状态，自动为命令行输出添加颜色，以提高可读性和用户体验。

## 二、命令使用

### 命令格式

```bash
git config --global color.ui auto
```

### 命令示例

```bash
$ git config --global color.ui auto
```

## 三、输出说明

执行该命令后，没有直接的输出信息。但后续使用 Git 命令时，命令行输出将会自动应用颜色，例如分支名称、提交信息等会以不同颜色显示。

## 四、注意事项

### 1. 全局配置

`--global` 参数表示该配置将应用于当前用户的所有 Git 仓库。如果只想对当前仓库生效，可以省略 `--global` 参数。

### 2. 取消彩色输出

如果需要取消彩色输出，可以使用以下命令：

```bash
git config --global color.ui false
```

### 3. 颜色自定义

如果需要自定义 Git 输出的颜色，可以通过 `git config` 命令设置具体的颜色选项。例如：

```bash
git config --global color.status.added green
```

该命令会将 `git status` 命令中新增文件的颜色设置为绿色。