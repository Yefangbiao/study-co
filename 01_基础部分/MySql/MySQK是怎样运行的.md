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