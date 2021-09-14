# 大纲

+ mysql基础(doing)
+ mysql进阶1:MySQL是怎样运行的 从根儿上理解MySQL
+ mysql进阶2:MySQL技术内幕:InnoDB存储引擎
+ mysql进阶3:高性能MySQL

# MySql

# 第1章 了解SQL

略

# 第2章 MySql简介

## 2.1 什么是MySQL

数据的所有存储,检索,管理和处理实际上都是由数据库软件--DBMS(数据库管理系统)来完成的.MySQL是一种DBMS,即一种数据库软件.

### 2.1.1 客户机-服务器软件

DBMS可以分为两类:一类基于共享文件系统的DBMS,另一类为基于客户机-服务器的DBMS.前者(包括Microsoft Access和FileMaker)用于桌面用途,通常不用于高端或更关键的应用.

MySql,Oracle以及Microsoft SQL Server等数据库是基于客户机-服务器的数据库.客户机-服务器应用分为两个不同的部分.服务器部分是负责所有数据访问和处理一个软件.这个软件运行在称为**数据库服务器**的计算机上.

与数据文件打交道的只有服务器软件.关于数据/数据添加/删除和数据更新的所有请求都由服务器软件完成.这些请求或更改来自运行客户机软件的计算机.**客户机**是与用户打交道的软件.用户发出的请求通过客户机软件经由网络提交该请求给服务器软件.服务器软件处理这个请求,然后将结果返回给客户机软件.

## 2.2 MySQL工具

### 2.2.1 mysql命令行实用程序

> mac使用mysql登录示例

`mysql -u root -p` ; 然后输入密码

+ 命令使用`;`或者`\g`结束,换句话说,仅按Enter不执行命令.
+ 输入`quit`和`exit`推出命令行实用程序

### 2.2.2 mysql图形化界面(推荐)

+ Navicat
+ SQLyog

## 2.3 小结

本章介绍了什么是MySql,引入了几个客户机实用程序.



# 第3章 使用MySQL

## 3.1 连接

为了连接到MySQL,需要以下信息:

+ 主机名(计算机名)----如果连接到本地MySQL服务器,为`localhost`
+ 端口(如果使用默认端口3306之外的端口)
+ 一个合法的用户名
+ 用户口令(如果需要)

在连接之后,你就可以访问你的登录名能够访问的任意数据库和表了.

## 3.2 选择数据库

最初连接到MySQL的时候,没有任何数据库可以打开供你使用.在你执行任意数据库操作之前,需要选择一个数据库.可以使用`USE`关键字.

> 关键字(key word) 作为MySQL语言组成部分的一个保留字.绝对不要用关键字命名一个表或列.

例如,为了使用`crashcourse`数据库,应该在命令行输入以下内容

```mysql
> use crashcourse;
> Database changed
```

分析:`USE`并不返回任何结果,而是依赖使用的客户机.

记住,必须先使用`USE`打开数据库,才能读取其中的数据.

## 3.3 了解数据库和表

如果你不知道可以使用的数据库名怎么办?

数据库/表/列/用户/权限等信息被存储在数据库和表中(MySQL使用MySQL来存储这些信息). 不过,内部的表一般不直接访问.可以使用`SHOW`命令显示这些信息.

看下面的例子:

```mysql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| crashcourse        |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```

分析:`SHOW databases;`返回可用数据库的一个列表



为了获得一个数据库内表的列表,需要使用`SHOW TABLES;`如下所示:

```mysql
mysql> SHOW TABLES;
+-----------------------+
| Tables_in_crashcourse |
+-----------------------+
| customers             |
| orderitems            |
| orders                |
| productnotes          |
| products              |
| vendors               |
+-----------------------+
6 rows in set (0.00 sec)
```

分析:`SHOW TABLES`返回当前可用数据库内可用表的列表.



`SHOW`也可以用来显示表列.

```mysql
mysql> SHOW COLUMNS FROM customers;
+--------------+-----------+------+-----+---------+----------------+
| Field        | Type      | Null | Key | Default | Extra          |
+--------------+-----------+------+-----+---------+----------------+
| cust_id      | int(11)   | NO   | PRI | NULL    | auto_increment |
| cust_name    | char(50)  | NO   |     | NULL    |                |
| cust_address | char(50)  | YES  |     | NULL    |                |
| cust_city    | char(50)  | YES  |     | NULL    |                |
| cust_state   | char(5)   | YES  |     | NULL    |                |
| cust_zip     | char(10)  | YES  |     | NULL    |                |
| cust_country | char(50)  | YES  |     | NULL    |                |
| cust_contact | char(50)  | YES  |     | NULL    |                |
| cust_email   | char(255) | YES  |     | NULL    |                |
+--------------+-----------+------+-----+---------+----------------+
9 rows in set (0.00 sec)
```

分析:`SHOW COLUMNNS`要求给出一个表明(这个例子中的FROM customers),它对每个字段返回一行,行中包括段名/数据类型/是否允许NULL/键信息/默认值以及其他信息(例如字段cust_id的auto_increment)

> 什么是自动增量?

某些列需要唯一值.例如,订单编号/雇员ID或顾客ID.在每个行添加到表中的时候,MySQL可以自动为每个行分配下一个可用编号,不用添加的时候手动分配唯一值.如果需要,则必须在用`CREATE`语句创建表的时候作为表定义的组成部分.

> DESCRIBE语句

MySQL支持用`DESCRIBE`作为`SHOW COLUMNS FROM`的一种快捷方式.换句话说,`DESCRIBE  customers;`是`SHOW COLUMNS FROM customers;`的一种快捷方式



一些其他的`SHOW`语句还有

+ `SHOW STATUS;`:用于显示广泛的服务器状态信息;
+ `SHOW CRRATE DATABASE {{数据库名}};`和`SHOW CREATE TABLE {{表名}};`:分别用来显示创建特定数据库或表的MySQL语句
+ `SHOW GRANTS;`:用来显示授予用户的安全权限.
+ `SHOW ERRORS;`和`SHOW WARNINGS;`用来显示服务器错误或警告消息.

> 进一步了解SHOW

在mysql命令行实用程序中,执行命令`HELP SHOW;`显示允许的`SHOW`语句.

## 3.4 小结

本章介绍了如何连接和登录MySQL,如何使用USE选择数据库,如何使用SHOW查看MySQL数据库/表和内部信息.

# 第4章 检索数据

本章介绍如何使用SELECT语句从表中检索一个或多个数据列.

## 4.1 SELECT语句

用途是从一个或多个表中检索信息.

为了使用`SELECT`检索表数据,必须至少给出两条信息--想选择什么,以及从什么地方选择.

### 4.2检索单个列

我们从简单的SQL SELECT开始,如下所示:

```mysql
mysql> select prod_name from products;
+----------------+
| prod_name      |
+----------------+
| .5 ton anvil   |
| 1 ton anvil    |
| 2 ton anvil    |
| Detonator      |
| Bird seed      |
| Carrots        |
| Fuses          |
| JetPack 1000   |
| JetPack 2000   |
| Oil can        |
| Safe           |
| Sling          |
| TNT (1 stick)  |
| TNT (5 sticks) |
+----------------+
14 rows in set (0.00 sec)
```

分析:上述语句利用`select`从products表中检索一个叫prod_name的列.所需的列名在SELECT关键字之后给出.`FROM`关键字支出从其中检索数据的表名.

> 未排序数据

返回的顺序可能是不一样的,因为没有排序.只要返回相同数目的行,就是正常的

> SQL语句和大小写

SQL语句不区分大小写,因此`SELECT`和`select`是相同的.同样,写成`Select`也没有关系.

## 4.3 检索多个列

要想从一个表中检索多个列,使用相同的`SELECT`语句.唯一的不同是必须在`SELECT`关键字后给出多个列名,列名之间必须用逗号分隔.

下列`SELECT`语句从products表中选择3列.

```mysql
mysql> select prod_id,prod_name,prod_price from products;
+---------+----------------+------------+
| prod_id | prod_name      | prod_price |
+---------+----------------+------------+
| ANV01   | .5 ton anvil   |       5.99 |
| ANV02   | 1 ton anvil    |       9.99 |
| ANV03   | 2 ton anvil    |      14.99 |
| DTNTR   | Detonator      |      13.00 |
| FB      | Bird seed      |      10.00 |
| FC      | Carrots        |       2.50 |
| FU1     | Fuses          |       3.42 |
| JP1000  | JetPack 1000   |      35.00 |
| JP2000  | JetPack 2000   |      55.00 |
| OL1     | Oil can        |       8.99 |
| SAFE    | Safe           |      50.00 |
| SLING   | Sling          |       4.49 |
| TNT1    | TNT (1 stick)  |       2.50 |
| TNT2    | TNT (5 sticks) |      10.00 |
+---------+----------------+------------+
14 rows in set (0.00 sec)
```

分析:与前一个例子一样,这条语句使用`SELECT`语句从表products中选择数据.

## 4.4 检索所有列

`SELECT`语句可以检索所有的列而不必逐个列出它们.可使用`*`通配符来达到.

```mysql
mysql> select * from products;
```

分析:给定一个通配符`*`,则返回表中所有列.

## 4.5 检索不同的行

`SELECT`返回所有匹配的行.如果不想要每个值都出现,怎么办?

可以使用`DISTINCT`关键字,这个关键字指示MySQL返回不同的值.

```mysql
mysql> select distinct vend_id from products;
+---------+
| vend_id |
+---------+
|    1001 |
|    1002 |
|    1003 |
|    1005 |
+---------+
4 rows in set (0.00 sec)
```

分析:`DISTINCT`告诉MySQL只返回不同(唯一)的vend_id行.如果使用`DISTINCT`关键字,必须直接放在列名前面.

## 4.6 限制结果

`SELECT`语句返回所有匹配的行,为了返回第一行或前几行,可以使用`LIMIT`子句.下面举一个栗子:

```mysql
mysql> select prod_name from products limit 5;
+--------------+
| prod_name    |
+--------------+
| .5 ton anvil |
| 1 ton anvil  |
| 2 ton anvil  |
| Detonator    |
| Bird seed    |
+--------------+
5 rows in set (0.00 sec)
```

分析:这个语句使用`SELECT`检索单个列,`LIMIT 5`指示MySQL返回不多于5行.

为了得出下一个5行,指定要检索的开始行和行数.如下所示

```mysql
mysql> select prod_name from products limit 5,5;
+--------------+
| prod_name    |
+--------------+
| Carrots      |
| Fuses        |
| JetPack 1000 |
| JetPack 2000 |
| Oil can      |
+--------------+
5 rows in set (0.00 sec)
```

分析:`Limit5,5`指示MySQL返回从行5开始的5行.第一个数为开始的位置,第二个数为要检索的行数.

> 行0 第一行不是0而是1

> 行数不够的时候,将只返回能返回的那些行.

> LIMIT语法:MySQL支持LIMIT 4 OFFSET 3,代表从行3开始取4行

## 4.7 使用完全限定的表名

可以受用完全限制的名字来引用列(同时使用表名和列字)

```mysql
mysql> select products.prod_name from products;
```

## 4.8 小结

本章学习了如何使用SQL的SELECT语句来检索单个表列/多个表列一级所有表列.下一章讲如何排序检索出来的数据.

# 第5章 排序检索数据

## 5.1 排序数据

如果不使用排序,数据一般将以它在底层表中出现的顺序显示.如果对数据进行更新和删除,也会影响顺序.

> 子句

SQL语句由子句构成,有些子句是必需的,而有的是可选的.一个子句通常由一个关键字提供的数据组成.

为了明确地排序检索的数据,可以使用`ORDER BY`子句.`ORDER BY`子句取一个或者多个列的名字,然后对输出进行排序.

```mysql
mysql> select prod_name from products order by prod_name;
+----------------+
| prod_name      |
+----------------+
| .5 ton anvil   |
| 1 ton anvil    |
| 2 ton anvil    |
| Bird seed      |
| Carrots        |
| Detonator      |
| Fuses          |
| JetPack 1000   |
| JetPack 2000   |
| Oil can        |
| Safe           |
| Sling          |
| TNT (1 stick)  |
| TNT (5 sticks) |
+----------------+
14 rows in set (0.00 sec)
```

分析:这条语句除了指示MySQL对prod_name以字母顺序排序之外,与前面的语句相同.

## 5.2 按照多个列进行排序

通常不止一个列进行数据排序.为了对多个列进行排序,需要指定列名,然后用逗号分开即可.

```mysql
mysql> select prod_id,prod_price,prod_name from products order by prod_price,prod_name;
+---------+------------+----------------+
| prod_id | prod_price | prod_name      |
+---------+------------+----------------+
| FC      |       2.50 | Carrots        |
| TNT1    |       2.50 | TNT (1 stick)  |
| FU1     |       3.42 | Fuses          |
| SLING   |       4.49 | Sling          |
| ANV01   |       5.99 | .5 ton anvil   |
| OL1     |       8.99 | Oil can        |
| ANV02   |       9.99 | 1 ton anvil    |
| FB      |      10.00 | Bird seed      |
| TNT2    |      10.00 | TNT (5 sticks) |
| DTNTR   |      13.00 | Detonator      |
| ANV03   |      14.99 | 2 ton anvil    |
| JP1000  |      35.00 | JetPack 1000   |
| SAFE    |      50.00 | Safe           |
| JP2000  |      55.00 | JetPack 2000   |
+---------+------------+----------------+
14 rows in set (0.00 sec)
```

分析:在排序多个列的时候,具有相同的prod_price时才会对prod_name进行排序.

## 5.3 指定排序方向

数据排序不限于升序排序(默认),也可以降序排序,为了降序排序,需要指定`DESC`关键字.

```mysql
mysql> select prod_id,prod_price,prod_name from products order by prod_price DESC,prod_name;
+---------+------------+----------------+
| prod_id | prod_price | prod_name      |
+---------+------------+----------------+
| JP2000  |      55.00 | JetPack 2000   |
| SAFE    |      50.00 | Safe           |
| JP1000  |      35.00 | JetPack 1000   |
| ANV03   |      14.99 | 2 ton anvil    |
| DTNTR   |      13.00 | Detonator      |
| FB      |      10.00 | Bird seed      |
| TNT2    |      10.00 | TNT (5 sticks) |
| ANV02   |       9.99 | 1 ton anvil    |
| OL1     |       8.99 | Oil can        |
| ANV01   |       5.99 | .5 ton anvil   |
| SLING   |       4.49 | Sling          |
| FU1     |       3.42 | Fuses          |
| FC      |       2.50 | Carrots        |
| TNT1    |       2.50 | TNT (1 stick)  |
+---------+------------+----------------+
14 rows in set (0.00 sec)
```

分析:`DESC`关键字只应用到直接位于其前面的列名.在上面,只对prod_price指定`DESC`,而prod_name仍是升序排序



使用`ORDER BY`和`LIMIT`组合,能找出一个列中最高或者最低的值.

```mysql
mysql> select prod_price from products order by prod_price desc limit 1;
+------------+
| prod_price |
+------------+
|      55.00 |
+------------+
1 row in set (0.00 sec)
```

> ORDER BY子句的位置

使用`ORDER BY`子句的时候,应该保证位于`FROM`子句之后.如果使用`LIMIT`,必须位于`ORDER BY`之后.

## 5.4 小结

本章学习了如何使用`SELECT`语句的`ORDER BY`子句对检索出的数据进行排序.这个子句必须是`SELECT`中最后一条子句.可根据需要,用它在一个或多个列上数据进行排序.
