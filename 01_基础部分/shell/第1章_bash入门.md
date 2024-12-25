## 为什么是bash

bash 是一种 shell，也就是命令解释器。bash（或者说任何 shell）的主要目的是让用户可以同计算机操作系统交互，以便完成想做的任务。这通常涉及运行程序，因此 shell 会接受你输入的命令，判断要用到的程序，然后执行命令来启动程序。你还会碰到一些需要执行一系列操作的任务，这些操作要么是重复性的，要么非常复杂。shell 编程（通常称为 **shell 脚本编程**）允许你对此类任务进行自动化，以实现易用性、可靠性以及可重现性。



## 查找并运行命令

你需要在 bash 下查找并运行特定的命令。

### 解决方案

可以试试 `type`、`which`、`apropos`、`locate`、`slocate`、`find` 和 `ls` 命令。

+ bash 会在环境变量 PATH 中保留一个用于查找命令的目录列表。内建命令 `type` 会在环境（包括别名、关键字、函数、内建命令、`$PATH` 中的目录以及命令散列表）中搜索匹配其参数的可执行文件并显示匹配结果的类型和位置。
+ `which` 命令与 `type` 类似，但它只搜索 `$PATH`（以及 csh 别名）。
+ 几乎所有的命令都自带某种形式的用法帮助。通常采用的是称为`手册页`（manpage）的在线文档，其中 man 是 manual（手册）的简写。可以使用 `man` 命令访问这些手册页，`man ls` 会显示 `ls` 命令的相关文档。很多程序还有内建的帮助机制，通过 `-h` 或 `--help` 这样的“帮助”选项就能使用。
+ 但如果你不知道或忘记了命令名，该怎么办呢？`apropos` 命令可以根据所提供的正则表达式参数搜索手册页名称及描述。

```shell
$ apropos music
cms (4) - Creative Music System device driver

$ man -k music
cms (4) - Creative Music System device driver
```

+ `locate` 和 `slocate` 通过查询系统数据库文件（通常由调度程序 cron 运行的作业负责编译和更新）来查找文件或命令，几乎立刻就能得到结果。在多数 Linux 系统中，`locate` 是指向 `slocate` 的符号链接；在其他系统中，两者可能是不同的程序，也可能根本就没有 `slocate`。

```shell
$ locate apropos
/usr/bin/apropos
/usr/share/man/de/man1/apropos.1.gz
/usr/share/man/es/man1/apropos.1.gz
/usr/share/man/it/man1/apropos.1.gz
/usr/share/man/ja/man1/apropos.1.gz
/usr/share/man/man1/apropos.1.gz
```

## 获取文件相关信息

使用 `ls`、`stat`、`file` 或 `find` 命令：

```shell
$ touch /tmp/sample_file

$ ls /tmp/sample_file
/tmp/sample_file

$ ls -l /tmp/sample_file
-rw-r--r-- 1 jp         jp            0 Dec 18 15:03 /tmp/sample_file
$ stat -x /tmp/sample_file
File: "/tmp/sample_file"
Size: 0           Blocks: 0        IO Block: 4096   Regular File
Device: 303h/771d Inode:  2310201    Links: 1
Access: (0644/-rw-r--r--) Uid: (  501/      jp)   Gid: ( 501/        jp)
Access: Sun Dec 18 15:03:35 2005
Modify: Sun Dec 18 15:03:35 2005
Change: Sun Dec 18 15:03:42 2005

$ file /tmp/sample_file
/tmp/sample_file: empty

$ file -b /tmp/sample_file
empty

$ echo '#!/bin/bash -' > /tmp/sample_file

$ file /tmp/sample_file
/tmp/sample_file: Bourne-Again shell script text executable

$ file -b /tmp/sample_file
Bourne-Again shell script text executable
```

