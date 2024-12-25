## 统计参数数量

你想知道调用脚本时使用了多少个参数。

使用 shell 内建变量 `$#`。例 5-4 展示了一个严格要求 3 个参数的脚本。

```bash
#!/usr/bin/env bash
# 实例文件：check_arg_count
#
# 检查正确的参数数量：
# 使用下列语法或者：if [ $# -lt 3 ]
if (( $# < 3 ))
then
    printf "%b" "Error. Not enough arguments.\n" >&2
    printf "%b" "usage: myscript file1 op file2\n" >&2
    exit 1
elif (( $# > 3 ))
then
    printf "%b" "Error. Too many arguments.\n" >&2
    printf "%b" "usage: myscript file1 op file2\n" >&2
    exit 2
else
    printf "%b" "Argument count correct. Proceeding...\n"
fi
```

