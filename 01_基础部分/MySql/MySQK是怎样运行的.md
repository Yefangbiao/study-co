# MySQL是怎样运行的

## 第1章-初识MySQL

假设已经把mysql的bin目录添加到了环境变量

## 1.1 客户端与服务器连接的过程

运行中的服务器程序和客户端本质都是一个计算机上的进程，客户端向服务器发送请求并且得到响应实际是一个进程间通信的过程

+ TCP/IP

真实环境中，数据库服务器进程和客户端进程可能运行在不同的主机中，它们之间必须通过网络进行通信。MySQL采用TCP作为服务器和客户端之间的网络通信协议。网络中的其他进程可以通过IP地址+端口号的方式与这个进程建立连接，这样进程之间就可以进行网络通信了。

MySQL在服务启动的时候默认申请3306端口号。

如果3306已经被占用，可以通过`-p`参数来指定端口号。比如

```mysql
mysqld -P3307
```

+ 命名管道和共享内存
  + 使用命名管道和共享内存进行通信，需要在启动服务器程序的命令加上`--enable-named-pipe`参数。然后在启动客户端的命令加上`--pipe`或者`--protocol=pipe`
  + 使用共享内存进行通信，需要在服务端加上`--shared-memory`参数。在客户端加上`--protocol=memory` 
+ UNIX域套接字

如果服务器进程和客户端进程都运行在**类UNIX的同一台机器上**。可以使用UNIX域套接字进行通信。如果启动客户端的时候没有指定主机名，或者指定的主机名为localhost，或者指定了`--protocol=socket`。那么服务器和客户端程序就可以通过UNIX套接字进行通信了。

MySQL服务器默认UNIX域套接字名称为`\tmp\mysql.sock`。如果想改变这个名称，可以指定socket参数。`mysqld --socket=/tmp/a.txt`。客户端需要指定连接的UNIX域套接字文件名称:`mysql -hlocalhost -uroot --socket=/tmp/a.txt`

## 1.2 服务器处理客户端请求

无论服务器和客户端采用什么方式通信。最后都是客户端进程向服务器发送一段文本(MySQL语句)，服务器进程处理后再向客户端进程返回一段文本（处理结果）。大致可以如下所示

![image-20210922214253684](MySQK是怎样运行的.assets/image-20210922214253684.png)

+ 连接管理：当有一个客户端进程连接到服务器进程的时候，服务器会创建一个线程与客户端进行交互。当客户端断开连接的时候，服务器不会理解销毁线程，而是缓存起来，当有一个新的客户端进行连接的时候，把这个缓存的线程分配给新客户端

+ 解析与优化：

  + 查询缓存
    + MySQL会把刚刚处理过的查询请求和结果缓存起来。如果下一次有同样的请求，直接从缓存中查找结果即可。
    + 如果有两个查询清幽有任何字符上的不同（空格、注释、大小写），都会导致缓存不命中。
    + 有些请求不会缓存，例如mysql、information_schema、performance_schema数据库中的表的请求。还包括一些系统函数例如`NOW()`
    + 缓存失效：当表的结构或者数据被修改，则与该表的所有查询缓存都将无效并且从查询缓存中删除。

  > MySQL5.7.20开始，不推荐使用查询缓存，在MySQL8.0中被删除

  + 语法解析：客户端发送的请求是一段文本，这时候MySQL服务器需要对这段文本进行分析，判断语法是否正确，然后从文本中将要查询的表、各种查询条件提取出来放到MySQL服务器内部使用的一些数据结构。
  + 查询优化：MySQL会对我们的语句做一些优化，比如外连接转换为内连接、表达式简化、子查询转换为连接

+ 存储引擎：MySQL把数据存储和提取操作都封装到了一个名为存储引擎的模块中。

MySQL服务器处理请求的过程简单划分为server层和存储引擎层。连接管理、查询缓存、语法解析、查询优化是server层。真正存取数据的层称为存储引擎层。

server层和存储引擎层交互的时候，一般以记录为单位。以SELECT为例。server层根据执行计划先向存储引擎层取一条记录，判断是否符合条件，如果符合，发送给客户端，否则跳过。

## 1.3 常用的存储引擎

| Feature                                | MyISAM       | Memory           | InnoDB       | Archive      | NDB          |
| :------------------------------------- | :----------- | :--------------- | :----------- | :----------- | :----------- |
| B-tree indexes                         | Yes          | Yes              | Yes          | No           | No           |
| Backup/point-in-time recovery (note 1) | Yes          | Yes              | Yes          | Yes          | Yes          |
| Cluster database support               | No           | No               | No           | No           | Yes          |
| Clustered indexes                      | No           | No               | Yes          | No           | No           |
| Compressed data                        | Yes (note 2) | No               | Yes          | Yes          | No           |
| Data caches                            | No           | N/A              | Yes          | No           | Yes          |
| Encrypted data                         | Yes (note 3) | Yes (note 3)     | Yes (note 4) | Yes (note 3) | Yes (note 3) |
| Foreign key support                    | No           | No               | Yes          | No           | Yes (note 5) |
| Full-text search indexes               | Yes          | No               | Yes (note 6) | No           | No           |
| Geospatial data type support           | Yes          | No               | Yes          | Yes          | Yes          |
| Geospatial indexing support            | Yes          | No               | Yes (note 7) | No           | No           |
| Hash indexes                           | No           | Yes              | No (note 8)  | No           | Yes          |
| Index caches                           | Yes          | N/A              | Yes          | No           | Yes          |
| Locking granularity                    | Table        | Table            | Row          | Row          | Row          |
| MVCC                                   | No           | No               | Yes          | No           | No           |
| Replication support (note 1)           | Yes          | Limited (note 9) | Yes          | Yes          | Yes          |
| Storage limits                         | 256TB        | RAM              | 64TB         | None         | 384EB        |
| T-tree indexes                         | No           | No               | No           | No           | Yes          |
| Transactions                           | No           | No               | Yes          | No           | Yes          |
| Update statistics for data dictionary  | Yes          | Yes              | Yes          | Yes          | Yes          |

## 第2章-启动选项和系统变量

略

## 第3章-字符集和比较规则

 ## 3.1 字符集

### 3.2.1 MySQL中的utf8和utf8mb4

+ utf8mb3：只使用1-3字节表示字符
+ Utf8mb4：正宗的UTF-8字符集，使用1-3字节表示字符

在MySQL中，utf8是utf8mb3的别名。

## 第4章-InnoDB记录存储结构

### 4.1 InnoDB页简介

InnoDB是一个将表中的数据存储到磁盘上的存储引擎。真正处理数据的过程发生在内存中，所以需要把磁盘中的数据加载到内存。如果处理写入或者修改请求，需要把内存的内容刷新到磁盘上。InnoDB将数据划分为若干页，以页作为磁盘和内存之间交互的单位。页大小一般为16KB。

### 4.2 InnoDB行格式

我们平时都是以记录为单位向表中插入数据的，这些记录在磁盘上的存放方式称为行格式或者记录格式。有四种不同的行格式，`COMPACT`、`REDUNDANT`、`DYNAMIC`和`COMPRESSED`

#### 4.2.1 指定行格式的语法

`CREATE TABLE {{表名}} (列的信息) ROW_FORMAT=行格式名称;`

`ALTER TABLE {{表名}} ROW_FORMAT=行格式名称;`

#### 4.3.2 COMPACT行格式

![image-20210922222624241](MySQK是怎样运行的.assets/image-20210922222624241.png)

+ 记录的额外信息：这部分是服务器为了管理记录不得不添加的一些信息。

  + 变长字段长度列表：MySQL支持变长字段，比如VARCHAR(M),VARBINARY(M)、各种TEXT类型、各种BLOB类型。变长字段存储数据大小不确定，所以存储真实数据的时候需要顺便把这些数据占用的字节数存储起来。

  > 在COMPACT中，所有变长字段的真实数据占用的字节数都在记录的开头位置。各变长字段的真实数据占用的字节数按照列的顺序逆序存放。**逆序存放**

  + NULL值列表:COMPAT把一条记录中值为NULL的列统一管理起来，存储到NULL值列表
    + 首先统计表中允许存储NULL的列
    + 如果表中没有存储NULL的列，则NULL不存在了。否则每一个NULL的列对应一个二进制位。二进制位1，代表值是NULL。二进制位0，代表值不为NULL。
    + NULL列必须用整个字节位表示，如果NULL数量不是整数个，最高位补0 
  + 记录头信息：固定5字节组成。

![image-20210922224339392](MySQK是怎样运行的.assets/image-20210922224339392.png)

+ + deleted_flag:标记是否被删除

  + min_rec_flag:B+树每层非叶子结点最小的目录项记录会添加该标记
  + n_owned:一个页面中记录会被分成若干组。第一个组的n_owned代表该小组中所有记录条数
  + heap_no:当前记录在页面堆中的相对位置
  + record_type:当前记录类型。0：普通，1：B+树非叶子结点目录项记录。2：infimum记录。3：supremum记录
  + next_record：下一条记录的相对位置

+ 记录的真实数据
  + MySQL会为每个记录默认添加一些列。
    + row_id:非必要，行ID，唯一标识一条记录
    + trx_id：必要，事务ID
    + roll_pointer：必要，回滚指针

InnoDB主键生成策略：优先使用用户自定义的主键，否则选一个一个UNIQUE键作为主键。否则默认添加一个row_id作为主键

#### 4.3.3 REDUNDANT行格式

![image-20210922225532207](MySQK是怎样运行的.assets/image-20210922225532207.png)

+ 字段长度偏移列表：逆序，使用偏移长度，例如`0C 06`即第一列0x06,第二列0x0c字节。
+ NULL值处理：对应偏移量的第一个比特位。如果是1，就是NULL。

#### 4.3.4 溢出列

如果一个字符串超过了页数，怎么办呢

在COMPACThe REDUNDANT中，对占用存储了非常多的列。只存储一部分，剩余数据存储在其他页中。记录真实数据用20字节指向这些页的地址

#### 4.3.5 DYNAMIC和COMPRESSED

+ DYNAMIC：类似于COMPACT。处理溢出的时候把所有数据存储到溢出页中。真实数据存储20字节大小的溢出页地址。
+ COMPRESSED使用压缩算法压缩。

## 第5章-InnoDB数据页结构

存放表中记录的那种类型的页。**索引(INDEX)页**。

### 5.1 数据页结构快览

![image-20210923211515572](MySQK是怎样运行的.assets/image-20210923211515572.png)

| 名称             | 中文名                     | 占用空间大小 | 简单描述               |
| ---------------- | -------------------------- | ------------ | ---------------------- |
| File Header      | 文件头部                   | 38字节       | 页的一些通用信息       |
| Page Header      | 页面头部                   | 56字节       | 数据页专有的一些信息   |
| Infimum+Supremum | 页面中的最小记录和最大记录 | 26字节       | 两个虚拟的记录         |
| User Records     | 用户记录                   | 不确定       | 用户存储的记录内容     |
| Free Space       | 空闲空间                   | 不确定       | 页中尚未使用的空间     |
| Page Directory   | 页目录                     | 不确定       | 页中某些记录的相对位置 |
| File Trailer     | 文件尾部                   | 8字节        | 校验页是否完整         |

### 5.2 记录在页中的存储

我们自己存储的记录会放在`User Records`部分。每当插入一条数据，会从`Free Space`部分申请一个记录大小的空间，并且将这个空间划分到`User Record`部分。

### 5.3 页目录

+ 将所有正常的记录（包括Infimum和Supremum）划分为几个组
+ 每个组最后一条记录（也是组内最大的那条记录），最后一条记录的`n_owned`表示组内一共有几条记录
+ 每个组最后一条记录在页面中的地址偏移量单独提取出来，按照顺序存储到靠近页尾部的地方。这个地方就是`Page Directory`

比如，现在又6条记录，会被分为两组

![image-20210923213838343](MySQK是怎样运行的.assets/image-20210923213838343.png)

+ 对于Infimum记录在的分组只能有一条记录
+ Supremum分组记录在1-8之间
+ 剩下的条数范围只能在4-8之间
  + 初始情况下，一个数据页中只有Infimum记录和Supremum。分属两个组。页目录中也只有两个槽。代表Infimum和Supremum在页中的地址偏移量
  + 每插入一条记录，都会在页目录中找到对应记录的主键值比待插入记录的主键值大并且差值最小的槽。然后把该槽对应的记录的`n_owned`加1，直到该组记录等于8个。
  + 当一个组记录等于8后，在插入一条记录，会分为两个组。其中一个4条记录，另一个5条记录。这个过程会在页目录中新增一个槽。

## 第6章-B+树索引

### 6.1 索引

![image-20210923215936847](MySQK是怎样运行的.assets/image-20210923215936847.png)

+ 目录项的`record_type`值是1
+ 目录项只有主键值和页的编号两个列。

多个目录页可以再组合

![image-20210923220156057](MySQK是怎样运行的.assets/image-20210923220156057.png)

1.联合索引

我们想让B+树按照c2和c3排列

+ 先把各个记录的页按照c2进行排列

+ 然后再采用c3进行排列

2.聚簇索引

+ 使用记录主键值进行记录和页的排序
+ 叶子结点存储了完整的用户记录

满足这两个条件称为聚簇索引

3.二级索引

聚簇索引只能搜索条件是主键的时候发挥作用。

我们可以多建几颗B+树。按照不同的排序规则排序

找到叶子结点后，叶子结点存储了主键信息，然后到聚簇索引去查找

## 第7章-B+树索引使用

略

## 第8章-MySQL的数据目录

如何确定MySQL中的数据目录

```mysql
SHOW VARIABLES LIKE 'datadir';
```

### 8.1 数据目录的结构

每次使用`CREATE DATABASE`的时候。MySQL会做两件事

+ 在数据目录下创建一个与数据库同名的子目录
+ 在子目录下创建一个`db.opt`的文件，包含了数据库的一些属性

### 8.2 表在文件系统中的表示

+ 保存表结构信息：`表名.frm`
+ 保存表数据：`表名.ibd`