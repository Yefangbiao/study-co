## 在shell脚本中执行算术操作

你需要在 shell 脚本中执行一些简单的算术操作。

用 `$(( ))` 或 `let` 进行整数运算。例如：

```bash
COUNT=$((COUNT + 5 + MAX * 2))
let COUNT+='5+MAX*2'
```

1. 赋值运算符两边不能出现空格
2. let语句中的每个单词本身就代表单独的算术表达式（在两个变量的两次赋值操作中，可以不用逗号）

## 条件分支

你想要检查参数数量是否正确并执行相应的操作。这得用到条件分支。

`bash` 中的 if 语句表面上和其他语言中的差不多。

```shell
if [ $# -lt 3 ]
then
    printf "%b" "Error. Not enough arguments.\n"
    printf "%b" "usage: myscript file1 op file2\n"
    exit 1
fi
```

或者：

```shell
if (( $# < 3 ))
then
    printf "%b" "Error. Not enough arguments.\n"
    printf "%b" "usage: myscript file1 op file2\n"
    exit 1
fi
```

以下是一个带有 elif（bash 中的 else-if）和 else 子句的完整 if 语句。

```shell
if (( $# < 3 ))
then
    printf "%b" "Error. Not enough arguments.\n"
    printf "%b" "usage: myscript file1 op file2\n"
    exit 1
elif (( $# > 3 ))
then
    printf "%b" "Error. Too many arguments.\n"
    printf "%b" "usage: myscript file1 op file2\n"
    exit 2
else
    printf "%b" "Argument count correct. Proceeding...\n"
fi
```

Linux中变量`#, @, 0, 1, 2, *,$$,$?`的含义

`$#` 是传给脚本的参数个数
`$0` 是脚本本身的名字
`$1` 是传递给该shell脚本的第一个参数
`$2` 是传递给该shell脚本的第二个参数
`$@` 是传给脚本的所有参数的列表
`$*` 是以一个单字符串显示所有向脚本传递的参数，与位置变量不同，参数可超过9个
`$$` 是脚本运行的当前进程ID号
`$?` 是显示最后命令的退出状态，0表示没有错误，其他表示有错误

## 测试文件特性

为了提高脚本的稳健性，你希望在读取输入文件前先检查该文件是否存在；另外，还想在写入输出文件前确认其是否具备写权限，在用 `cd` 切换目录前看看到底有没有这个目录。

在 `if` 语句的 `test` 命令部分使用各种文件特性测试。

```shell
#!/usr/bin/env bash
# 实例文件：checkfile
#
DIRPLACE=/tmp
INFILE=/home/yucca/amazing.data
OUTFILE=/home/yucca/more.results

if [ -d "$DIRPLACE" ]
then
    cd $DIRPLACE
    if [ -e "$INFILE" ]
    then
        if [ -w "$OUTFILE" ]
        then
            doscience < "$INFILE" >> "$OUTFILE"
        else
            echo "cannot write to $OUTFILE"
        fi
    else
        echo "cannot read from $INFILE"
    fi
else
    echo "cannot cd into $DIRPLACE"
fi
```

----

运算符	描述

-a  表示AND
-o  表示OR
-b	块设备文件（如 /dev/hda1）
-c	字符设备文件（如 /dev/tty）
-d	目录文件
-e	文件存在
-f	普通文件
-g	文件设置了 set-group-ID（setgid）位
-h	符号链接文件（等同于 -L）
-G	有效组 ID（effective group ID）拥有的文件
-k	文件设置了粘滞位
-L	符号链接文件（等同于 -h）
-N	文件自上次读取后被修改过
-O	有效用户 ID（effective user ID）拥有的文件
-p	具名管道文件
-r	可读文件
-s	文件大小不为空
-S	套接字文件
-u	文件设置了 set-user-ID（setuid）位
-w	可写文件
-x	可执行文件
-z    字符串长度为零为真

---

同一个语句中可以出现多个 AND/OR。你可能要用括号来获得正确的优先级，比如 `a and (b or c)`，但一定要记得在括号前加上反斜杠或将括号放进引号，以消除其特殊含义。不过，可别让整个表达式都出现在引号中，这会令其作为整体被当成是对空字符串的测试（参见 6.5 节）。以下这个测试更为复杂，其中的括号已经过正确转义。

```bash
if [ -r "$FN" -a \( -f "$FN" -o -p "$FN" \) ]
```

## 测试等量关系

你想要检查两个 shell 变量是否相等，但是存在两种测试运算符：`-eq` 和 `=`（或 `==`）。该用哪个呢？

你需要的比较类型决定了该用哪种运算符。如果是进行数值比较，可以使用 `-eq` 运算符；如果是进行字符串比较，则使用 `=`（或 `==`）运算符。

```shell
#!/usr/bin/env bash
# 实例文件：strvsnum
#
# 老生常谈的字符串与数值比较
#
VAR1=" 05 "
VAR2="5"
printf "%s" "do they -eq as equal? "
if [ "$VAR1" -eq "$VAR2" ]
then
    echo YES
else
    echo NO
fi

printf "%s" "do they = as equal? "
if [ "$VAR1" = "$VAR2" ]
then
    echo YES
else
    echo NO
fi
```

数值比较运算符	字符串比较运算符	含义
-lt	<	小于
-le	<=	小于或等于
-gt	>	大于
-ge	>=	大于或等于
-eq	=，==	等于
-ne	!=	不等于

## 用模式匹配进行测试

你不想对字符串进行字面匹配，而是想看看它是否符合某种模式。例如，想知道是否存在 JPEG 类型的文件。
在 `if` 语句中使用复合命令 `[[ ]]`，以便在等量运算符右侧启用 shell 风格的模式匹配：

```shell
if [[ "${MYFILENAME}" == *.jpg ]]
```

`[[ ]]` 语法不同于 test 命令的老形式 `[ ]`，它是一种较新的 bash 机制（2.01 版左右才出现）。能够在 `[ ]` 中使用的运算符也可用于 `[[ ]]`，但在后者中，等号是一种更为强大的字符串比较运算符。

## 用正则表达式进行测试

使用 `=~` 运算符进行正则表达式匹配。只要能够匹配到某个字符串，就可以在 shell 数组变量 `$BASH_REMATCH` 中找到模式中的各个部分所匹配到的内容。

## 循环一段时间

你希望在符合条件的情况下重复执行某些操作。

**对于算术条件，使用 `while` 循环：** 

```shell
while (( COUNT < MAX ))
do
    some stuff
    let COUNT++
done
```

**对于文件系统相关的条件：**

```shell
while [ -z "$LOCKFILE" ]
do
    some things
done
```

**对于读取输入：**

```shell
while read lineoftext
do
    process $lineoftext
done
```

