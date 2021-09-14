# 大纲

+ mysql基础(doing)
+ mysql进阶

# MySql

# 第一章 了解SQL

略

# 第二章 MySql简介

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



# 第三章 使用MySQL

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

