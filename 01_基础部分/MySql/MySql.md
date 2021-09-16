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

# 第6章 过滤数据

本章将讲授如何使用`SELECT`语句的`WHERE`子句指定搜索条件

## 6.1 使用WHERE子句

数据库表一般包含大量的数据,很少需要检索所有行,只检索需要的数据需要指定**搜索条件(search criteria)**,搜索条件也称为**过滤条件(filter condition)**

在`SELECT`语句中,数据根据`WHERE`子句中指定的搜索条件进行过滤.`WHERE`子句在表名(`FROM`子句)之后给出

```mysql
mysql> select prod_name,prod_price from products where prod_price = 2.5;
+---------------+------------+
| prod_name     | prod_price |
+---------------+------------+
| Carrots       |       2.50 |
| TNT (1 stick) |       2.50 |
+---------------+------------+
2 rows in set (0.00 sec)
```

分析:这条语句从products表中检索两个列,但不返回所有行,只返回prod_price值为2.50的行

> WHERE 子句的位置

在同时使用`ORDER BY`和`WHERE`子句时,应该让`ORDER BY`位于`WHERE`之后,否则会产生错误

## 6.2 WHERE子句操作符

MySQL支持下列的所有条件操作符

| 操作符  | 说明               |
| ------- | ------------------ |
| =       | 等于               |
| <>      | 不等于             |
| !=      | 不等于             |
| <       | 小于               |
| <=      | 小于等于           |
| >       | 大于               |
| >=      | 大于等于           |
| BETWEEN | 在指定的两个值之间 |

### 6.2.1 检查单个值

我们下面看一个例子

```mysql
mysql> select prod_name,prod_price from products where prod_name = 'fuses';
+-----------+------------+
| prod_name | prod_price |
+-----------+------------+
| Fuses     |       3.42 |
+-----------+------------+
1 row in set (0.00 sec)
```

分析:MySQL在执行匹配时默认不区分大小写,所以fuses与Fuses匹配.

### 6.2.2 不匹配检查

以下例子列出不是由供应商1003制造的所有产品.

```mysql
mysql> select vend_id,prod_name from products where vend_id <>1003;
+---------+--------------+
| vend_id | prod_name    |
+---------+--------------+
|    1001 | .5 ton anvil |
|    1001 | 1 ton anvil  |
|    1001 | 2 ton anvil  |
|    1002 | Fuses        |
|    1005 | JetPack 1000 |
|    1005 | JetPack 2000 |
|    1002 | Oil can      |
+---------+--------------+
7 rows in set (0.00 sec)
```

> 何时使用引号

如果仔细观察上述`WHERE`子句中使用的条件,会看到有的值在单引号内(\`fuses`).有的没有.

+ 单引号用来限制字符串,如果将值与串类型的列进行比较,需要限定引号.否则不用
+ 如果是数值类的进行比较,不需要引号

### 6.2.3 范围值检查

为了检查某个范围的值,可使用`BETWEEN`操作符.语法与其他的子句操作符稍有不同,因为它需要两个值.即范围的开始值和结束值.下面例子说明`BETWEEN`如何检索价格在5美元到10美元之间的所有产品.

```mysql
mysql> select prod_name,prod_price from products where prod_price between 5 and 10;
+----------------+------------+
| prod_name      | prod_price |
+----------------+------------+
| .5 ton anvil   |       5.99 |
| 1 ton anvil    |       9.99 |
| Bird seed      |      10.00 |
| Oil can        |       8.99 |
| TNT (5 sticks) |      10.00 |
+----------------+------------+
5 rows in set (0.00 sec)
```

分析:使用`BETWEEN`的时候,必须指定两个值--所需范围的低端值和高端值.这两个值必须用`AND`关键字分隔.`BETWEEN`匹配范围中所有的值,包括开始值和结束值.

### 6.2.4 空值检查

在创建表的时候,表设计人员可以指定其中的列是否可以不包含值.不包含值的时候,称为包含空值`NULL`.

> NULL 无值,与包含字段0/空字符串或者仅仅包含空格不同

`SELECT`语句有一个特殊的`WHERE`子句,可以用来检查具有`NULL`值的列.这个`WHERE`子句就是`IS NULL`子句

```mysql
mysql> select prod_name from products where prod_price is NULL;
Empty set (0.00 sec)
```

## 6.3 小结

本章介绍了如何用`SELECT`语句的`WHERE`子句过滤返回的数据.过滤条件有相等/不相等/大于/小于/BETWEEN一级NULL值等进行测试.

# 第7章 数据过滤

这章讲述如何组合`WHERE`子句来建立功能更强的更高级的搜索条件,我们还将学习如何使用`NOT`和`IN`操作符

## 7.1 组合WHERE子句

### 7.1.1 AND操作符

为了不止对一个列进行过滤,可以使用`AND`操作符给`WHERE`子句附加条件.

```mysql
mysql> select prod_id ,prod_price,prod_name from products where vend_id = 1003 and prod_price<=10;
+---------+------------+----------------+
| prod_id | prod_price | prod_name      |
+---------+------------+----------------+
| FB      |      10.00 | Bird seed      |
| FC      |       2.50 | Carrots        |
| SLING   |       4.49 | Sling          |
| TNT1    |       2.50 | TNT (1 stick)  |
| TNT2    |      10.00 | TNT (5 sticks) |
+---------+------------+----------------+
5 rows in set (0.00 sec)
```

分析:这个SQL语句检索由供应商1003 制造且价格小于等于10美元的所有产品的名称和价格.这个语句里面包含两个条件,并且使用`AND`关键字连接他们.

### 7.1.2 OR操作符

`OR`指示检索匹配任意条件的行.

```mysql
mysql> select prod_id,prod_price,prod_name from products where vend_id = 1003 or vend_id = 1004 order by prod_price desc;
+---------+------------+----------------+
| prod_id | prod_price | prod_name      |
+---------+------------+----------------+
| SAFE    |      50.00 | Safe           |
| DTNTR   |      13.00 | Detonator      |
| FB      |      10.00 | Bird seed      |
| TNT2    |      10.00 | TNT (5 sticks) |
| SLING   |       4.49 | Sling          |
| FC      |       2.50 | Carrots        |
| TNT1    |       2.50 | TNT (1 stick)  |
+---------+------------+----------------+
7 rows in set (0.00 sec)
```

分析:SQL语句检索任一个符合条件的供应商的信息(1003和1004)

### 7.1.3 计算次序

`WHERE`可包含任意数目的`AND`和`OR`操作符.允许两者结合进行更复杂和更高级的过滤.

但,组合`AND`和`OR`也带来了一些问题.假如需要列出价格为10美元(含)以上且由1002或1003制造的所有产品,下面是一个错误的示例

```mysql
mysql> select prod_name,prod_price from products where vend_id=1002 or vend_id=1003 and prod_price >=10;
+----------------+------------+
| prod_name      | prod_price |
+----------------+------------+
| Detonator      |      13.00 |
| Bird seed      |      10.00 |
| Fuses          |       3.42 |
| Oil can        |       8.99 |
| Safe           |      50.00 |
| TNT (5 sticks) |      10.00 |
+----------------+------------+
6 rows in set (0.00 sec)


```

分析:有两行价格小于10美元.这是一个错误的返回.

SQL在处理`OR`之前,优先处理`AND`操作符.当SQL看到上述`WHERE`子句的时候,理解为由供应商1003制造的10美元以上的商品,或由供应商1002制造的任何产品,而不关其价格如何.

解决:使用圆括号明确分组相应的操作符.看下面的正确示例

```mysql
mysql> select prod_name,prod_price from products where (vend_id=1002 or vend_id=1003) and prod_price >=10;
+----------------+------------+
| prod_name      | prod_price |
+----------------+------------+
| Detonator      |      13.00 |
| Bird seed      |      10.00 |
| Safe           |      50.00 |
| TNT (5 sticks) |      10.00 |
+----------------+------------+
4 rows in set (0.00 sec)
```

下面的才是正确的结果.

## 7.2 IN操作符

**圆括号**在`WHERE`子句中还有另外一种用法.`IN`操作符用来指定条件范围,范围中的每个条件都可以进行匹配.`IN`取合法值的由逗号分隔的清单,全都在圆括号中.

```mysql
mysql> select prod_name,prod_price from products where vend_id in (1002,1003) order by prod_name;
+----------------+------------+
| prod_name      | prod_price |
+----------------+------------+
| Bird seed      |      10.00 |
| Carrots        |       2.50 |
| Detonator      |      13.00 |
| Fuses          |       3.42 |
| Oil can        |       8.99 |
| Safe           |      50.00 |
| Sling          |       4.49 |
| TNT (1 stick)  |       2.50 |
| TNT (5 sticks) |      10.00 |
+----------------+------------+
9 rows in set (0.00 sec)
```

这个语句检索供应商1002和1003制造的所有产品.`IN`操作符后跟由逗号分隔的合法值清单,整个清单必须在圆括号中.

为什么使用`IN`操作符?

+ 使用长的合法选项清单时,`IN`操作符的语法更清楚且更直观
+ 使用`IN`时,计算的次序更容易管理
+ `IN`操作符一般比`OR`操作符清单执行更快
+ `IN`最大优点是可以包含其他`SELECT`语句,能更动态地建立`WHERE`语句

## 7.3 NOT操作符

`WHERE`子句中的`NOT`操作符有且只有一个功能,那就是否定它之后所跟的任何条件

下面的例子说明`NOT`的使用.列出除了1002和1003之外的所有供应商的产品

```mysql
mysql> select prod_name,prod_price from products where vend_id not in (1002,1003) order by prod_name;
+--------------+------------+
| prod_name    | prod_price |
+--------------+------------+
| .5 ton anvil |       5.99 |
| 1 ton anvil  |       9.99 |
| 2 ton anvil  |      14.99 |
| JetPack 1000 |      35.00 |
| JetPack 2000 |      55.00 |
+--------------+------------+
5 rows in set (0.00 sec)
```

> MySQL中的NOT

MySQL支持使用`NOT`对`IN`,`BETWEEN`和`EXISTS`子句取反

## 7.4 小结

本章讲授如何使用`AND`和`OR`操作符组成`WHERE`子句,而且还讲授了如何明确地管理计算的次序,如何使用`IN`和`NOT`操作符

# 第8章 用通配符进行过滤

## 8.1 LIKE操作符

前面介绍的所有操作符都是针对已知值进行过滤的.利用通配符可以创建比较特定的数据的搜索模式.

> 通配符 用来匹配值的一部分的特殊字符

> 搜索模式 由字面值,通配符或两者组合构成的搜索条件

通配符本身实际是SQL的`WHERE`子句中有特殊含义的字符,SQL支持几种通配符.

使用通配符,必须使用`LIKE`操作符.`LIKE`指示MySQL.后跟的搜索模式利用通配符匹配而不是直接相等匹配进行比较.

### 8.1.1 百分号(%)通配符

最常使用的通配符是百分号`%`.在搜索串中,`%`表示任何字符出现**任意次数**.例如,为了找到词jet起头的产品,可以如何使用

```mysql
mysql> select prod_id, prod_name from products where prod_name like 'jet%';
+---------+--------------+
| prod_id | prod_name    |
+---------+--------------+
| JP1000  | JetPack 1000 |
| JP2000  | JetPack 2000 |
+---------+--------------+
2 rows in set (0.00 sec)
```

分析:`jet%`表示检索任意`jet`起头的词.`%`表示接受`jet`之后任意字符,不管它有多少字符.

> % 匹配0个,1个或多个字符

### 8.1.2 下划线(_)通配符

另一个有用的通配符是下划线`_`.下划线的用途和`%`一样,但`_`匹配单个字符而不是多个字符

```mysql
mysql> select prod_id, prod_name from products where prod_name like '_ ton anvil';
+---------+-------------+
| prod_id | prod_name   |
+---------+-------------+
| ANV02   | 1 ton anvil |
| ANV03   | 2 ton anvil |
+---------+-------------+
2 rows in set (0.00 sec)
```

与`%`不同,`_`总是匹配一个字符,不能多也不能少.

## 8.2 使用通配符的技巧

通配符一般比其他搜索所花的时间更长.

+ 不要过度使用通配符,如果其他操作符能达到相同目的,应该使用其他通配符
+ 确实需要通配符的时候,除非有绝对必要,否则不要用在搜索模式的开始处.通配符放在搜索模式的开始处,搜索是最慢的.

## 8.3 小结

本章介绍了什么是通配符以及如何在`WHERE`子句中使用通配符.

# 第9章 正则表达式

## 9.1 正则表达式介绍

略

## 9.2 使用MySQL正则表达式

这里只介绍正则表达式如何在MySQL中使用.

匹配`Jet`开头的prod_name

```mysql
mysql> select prod_name from products where prod_name regexp 'jet.*';
+--------------+
| prod_name    |
+--------------+
| JetPack 1000 |
| JetPack 2000 |
+--------------+
2 rows in set (0.00 sec)
```

## 9.3 总结

本章介绍了正则表达式的基础知识,学习如何在MySQL的`SELECT`语句中通过`REGEXP`关键字使用它们.

# 第10章 创建计算字段

## 10.1 计算字段

我们需要直接从数据库中检索出转换,计算或格式化过的数据;而不是检索出数据,然后在客户机应用程序或报告程序中重新格式化.

这就是计算字段发挥作用的所在了。与前面各章介绍过的列不同， 计算字段并不实际存在于数据库表中。计算字段是运行时在SELECT语句 内创建的。

## 10.2 拼接字段

为了说明如何使用计算字段,举一个创建两列组成的标题的简单例子.

vendors表中包含供应商名和位置信息.我们需要在供应商表中按照name(location)这样的格式列出供应商的位置.

此报表需要单个值,而表中数据存储在两个列vend_name和vend_country中括起来,这些东西没有明确存储在数据库表中

> 拼接(concatenate) 将值联结起来构成单个值

解决办法是把两个列拼接起来.在MySQL中,可以使用Concat()函数来拼接两个列

```mysql
mysql> select Concat(vend_name, '(',vend_country, ')') as 'name(location)' from vendors order by vend_name;
+------------------------+
| name(location)         |
+------------------------+
| ACME(USA)              |
| Anvils R Us(USA)       |
| Furball Inc.(USA)      |
| Jet Set(England)       |
| Jouets Et Ours(France) |
| LT Supplies(USA)       |
+------------------------+
6 rows in set (0.00 sec)
```

分析:Concat()拼接串，即把多个串连接起来形成一个较长的串。

Concat()需要一个或多个指定的串，各个串之间用逗号分隔。

上述的`SELECT`语句连接以下4个元素

+ 存储在vend_name列中的名字
+ 一个左圆括号
+ 存储在vend_country列中的国家
+ 包含一个右圆括号的串

> 使用别名

别名使用`AS`关键字赋值

## 10.3 使用算术计算

计算字段的另一常见用途是对检索出得数据进行算术计算.

orderitems表包含每个订单中的各项物品以及item_price列包含订单中每项物品的单价.而quantity包含数量.

```mysql
mysql> select prod_id, quantity, item_price, quantity*item_price as total from orderitems order by total desc;
+---------+----------+------------+---------+
| prod_id | quantity | item_price | total   |
+---------+----------+------------+---------+
| TNT2    |      100 |      10.00 | 1000.00 |
| FC      |       50 |       2.50 |  125.00 |
| ANV01   |       10 |       5.99 |   59.90 |
| JP2000  |        1 |      55.00 |   55.00 |
| TNT2    |        5 |      10.00 |   50.00 |
| ANV02   |        3 |       9.99 |   29.97 |
| ANV03   |        1 |      14.99 |   14.99 |
| FB      |        1 |      10.00 |   10.00 |
| FB      |        1 |      10.00 |   10.00 |
| OL1     |        1 |       8.99 |    8.99 |
| SLING   |        1 |       4.49 |    4.49 |
+---------+----------+------------+---------+
11 rows in set (0.00 sec)
```

MySQL支持如下的基本算术操作符

| 操作符 | 说明 |
| ------ | ---- |
| +      | 加   |
| -      | 减   |
| *      | 乘   |
| /      | 除   |

## 10.4 小结

本章介绍了计算字段以及如何创建计算字段.

# 第11章 使用数据处理函数

## 11.1 函数

SQL支持利用函数来处理数据,函数一般是在数据上执行的

## 11.2 使用函数

大多数SQL实现支持以下类型的函数。

+ 用于处理文本串(如删除或填充值,转换值为大写或小写)的文本函数
+ 用于在数值数据上进行算术操作(返回绝对值,进行代数运算)的数值函数
+ 用于处理日期和时间值并从这些值中提取特定成分(例如:返回两个日期之差,检查日期有效性等)的日期和时间函数
+ 返回DBMS正在使用的特殊信息(如返回用户登录信息,检查版本细节)的系统函数.

### 11.2.1 文本处理函数

使用`Upper()函数`

```mysql
mysql> select vend_name, Upper(vend_name) AS vend_name_upcase from vendors order by vend_name;
+----------------+------------------+
| vend_name      | vend_name_upcase |
+----------------+------------------+
| ACME           | ACME             |
| Anvils R Us    | ANVILS R US      |
| Furball Inc.   | FURBALL INC.     |
| Jet Set        | JET SET          |
| Jouets Et Ours | JOUETS ET OURS   |
| LT Supplies    | LT SUPPLIES      |
+----------------+------------------+
6 rows in set (0.00 sec)
```

分析:正如所见,`Upper()`将文本转换为大写.

下面列出了一些常用的文本处理函数:

| 函数        | 说明              |
| ----------- | ----------------- |
| Left()      | 返回串左边的字符  |
| Length()    | 返回串的长度      |
| Locate()    | 找出串的一个子串  |
| Lower()     | 将串转换为小写    |
| LTrim()     | 去掉串左边的空格  |
| Right()     | 返回串右边的字符  |
| RTrim()     | 去掉串右边的空格  |
| Soundex()   | 返回串的SOUNDEX值 |
| SubString() | 返回子串的字符    |
| Upper()     | 将串转换为大写    |

SOUNDEX需要做进一步的解释。SOUNDEX是一个将任何文 本串转换为描述其语音表示的字母数字模式的算法。

一般用不到

### 11.2.2 日期和时间处理函数

日期和时间采用相应的数据类型和特殊的格式存储，以便能快速和有效地排序或过滤，并且节省物理存储空间。

| 函数          | 说明                           |
| ------------- | ------------------------------ |
| AddDate()     | 增加一个日期(天,周等)          |
| AddTime()     | 增加一个时间(时、分等)         |
| CurDate()     | 返回当前日期                   |
| CurTime()     | 返回当前时间                   |
| Date()        | 返回日期时间的日期部分         |
| DateDiff()    | 计算两个日期之差               |
| Date_Add()    | 高度灵活的日期运算函数         |
| Date_Format() | 返回一个格式化的日期或时间串   |
| Day()         | 返回一个日期的天数部分         |
| DayOfWeek()   | 对于一个日期，返回对应的星期几 |
| Hour()        | 返回一个时间的小时部分         |
| Minute()      | 返回一个时间的分钟部分         |
| Month()       | 返回一个日期的月份部分         |
| Now()         | 返回当前日期和时间             |
| Second()      | 返回一个时间的秒部分           |
| Time()        | 返回一个日期时间的时间部分     |
| Year()        | 返回一个日期的年份部分         |

首先要注意MySQL使用的日期格式.无论什么时候指定一个日期.无论是插入还是更新还是使用`WHERE`进行过滤.日期格式都必须为`yyyy-mm-dd`.因此.2014年10月14日,给出的为2014-10-14

> 应该总是使用4位数字的年份

```mysql
mysql> select cust_id, order_num from orders where order_date='2005-09-01';
+---------+-----------+
| cust_id | order_num |
+---------+-----------+
|   10001 |     20005 |
+---------+-----------+
1 row in set (0.00 sec)
```

但是，使用`WHERE order_date = '2005-09-01'`可靠吗?order_ date的数据类型为datetime。这种类型存储日期及时间值。样例表中 的值全都具有时间值00:00:00，但实际中很可能并不总是这样。如果 用当前日期和时间存储订单日期(因此你不仅知道订单日期，还知道 下订单当天的时间)，怎么办?比如，存储的order_date值为 `2005-09-01 11:30:05`，则WHERE order_date = '2005-09-01'失败。 即使给出具有该日期的一行，也不会把它检索出来，因为WHERE匹配失 败。

解决办法是指示MySQL仅将给出的日期与列中的日期部分进行比 较，而不是将给出的日期与整个列值进行比较。为此，必须使用Date() 函数。

因此更可靠的语句是

```mysql
mysql> select cust_id, order_num from orders where Date(order_date)='2005-09-01';
```

不过，还有一种日期比较需要说明。如果你想检索出2005年9月下的 所有订单，怎么办?简单的相等测试不行，因为它也要匹配月份中的天 数。有几种解决办法，其中之一如下所示:

```mysql
mysql> select cust_id, order_num from orders where Date(order_date) BETWEEN '2005-09-01' AND '2005-09-30';
+---------+-----------+
| cust_id | order_num |
+---------+-----------+
|   10001 |     20005 |
|   10003 |     20006 |
|   10004 |     20007 |
+---------+-----------+
3 rows in set (0.00 sec)
```

还有另外一种如下所示:

```mysql
mysql> select cust_id, order_num from orders where Year(order_date) = 2005 AND Month(order_date) = 9;
+---------+-----------+
| cust_id | order_num |
+---------+-----------+
|   10001 |     20005 |
|   10003 |     20006 |
|   10004 |     20007 |
+---------+-----------+
3 rows in set (0.00 sec)
```

### 11.2.3 数值处理函数

数值处理函数用在处理数据.一般用在代数,三角函数,几何运算.

| 函数   | 说明                 |
| ------ | -------------------- |
| Abs()  | 返回一个函数的绝对值 |
| Cos()  | 返回一个角度的余弦   |
| Exp()  | 返回一个数的指数值   |
| Mod()  | 返回除操作的余数     |
| Pi()   | 返回圆周率           |
| Rand() | 返回一个随机数       |
| Sin()  | 返回一个角度的正弦   |
| Sqrt() | 返回一个数的平方根   |
| Tan()  | 返回一个角度的正切   |

## 11.3 小结

本章介绍了如何使用SQL的数据处理函数.

# 第12章 汇总数据

## 12.1 聚集函数

我们经常需要汇总数据而不用把它们实际检索出来，为此MySQL提 供了专门的函数。使用这些函数，MySQL查询可用于检索数据，以便分 析和报表生成。这种类型的检索例子有以下几种。

+ 确定表中行数
+ 获得表中行组的和
+ 找出表列的最大值,最小值和平均值

> 聚集函数 运行在行组上,计算和返回单个值的函数.

| 函数    | 说明             |
| ------- | ---------------- |
| AVG()   | 返回某列的平均值 |
| COUNT() | 返回某列的行数   |
| MAX()   | 返回某列的最大值 |
| MIN()   | 返回某列的最小值 |
| SUM()   | 返回某列值之和   |

### 12.1.1 AVG()函数

AVG()通过对表中行数计数并计算特定列值之和，求得该列的平均 值

```mysql
mysql> select AVG(prod_price) as avg_price from products;
+-----------+
| avg_price |
+-----------+
| 16.133571 |
+-----------+
1 row in set (0.00 sec)
```

### 12.1.2 COUNT()函数

COUNT()函数进行计数.可利用COUNT()确定表中行的数目或符合特定条件的行的数目.

COUNT()函数有两种使用方式.

+ 使用COUNT(*)对表中行的数目进行计算,不管表列中包含的是空(NULL)还是非空值
+ 使用COUNT(column)对特定列中具有值的行进行计数,胡烈NULL值

1.下面的例子返回customers表中客户的总数

```mysql
mysql> select COUNT(*) AS num_cust from customers;
+----------+
| num_cust |
+----------+
|        5 |
+----------+
1 row in set (0.00 sec)
```

2. 下列例子只对具有电子邮件地址的客户计数

```mysql
mysql> select COUNT(cust_email) as num_cust from customers;
+----------+
| num_cust |
+----------+
|        3 |
+----------+
1 row in set (0.00 sec)
```

### 12.1.3 MAX()函数

MAX()返回指定列中的最大值,要求指定列名

```mysql
mysql> select max(prod_price) as max_price from products;
+-----------+
| max_price |
+-----------+
|     55.00 |
+-----------+
1 row in set (0.00 sec)
```

> 对非数值数据使用MAX() 虽然MAX()一般用来找出最大的 数值或日期值，但MySQL允许将它用来返回任意列中的最大 值，包括返回文本列中的最大值。在用于文本数据时，如果数 据按相应的列排序，则MAX()返回最后一行。

> NULL值 MAX()函数忽略列值为NULL的行。

### 12.1.4 MIN()函数

MIN()功能恰好与MAX()相反,也要求指定列名

```mysql
mysql> select min(prod_price) as max_price from products;
+-----------+
| max_price |
+-----------+
|      2.50 |
+-----------+
1 row in set (0.00 sec)
```

> 对非数值数据使用MIN() MIN()函数与MAX()函数类似， MySQL允许将它用来返回任意列中的最小值，包括返回文本 列中的最小值。在用于文本数据时，如果数据按相应的列排序， 则MIN()返回最前面的行。

> NULL值 MIN()函数忽略列值为NULL的行

### 12.1.5 SUM()函数

SUM()用来返回指定列值的和(总计)。

下面举一个例子，orderitems表包含订单中实际的物品，每个物品 有相应的数量(quantity)。可如下检索所订购物品的总数(所有 quantity值之和):

```mysql
mysql> select SUM(quantity) as total from orderitems;
+-------+
| total |
+-------+
|   174 |
+-------+
1 row in set (0.00 sec)
```

SUM()也可以用来合计计算值。在下面的例子中，合计每项物品的

item_price*quantity，得出总的订单金额:

```mysql
mysql> select SUM(quantity*item_price) as total from orderitems;
+---------+
| total   |
+---------+
| 1368.34 |
+---------+
1 row in set (0.00 sec)
```

## 12.2 聚集不同值

以上5个聚集函数都可以如下使用:

+ 只包含不同的值,指定`DISTINCT`参数

下面的例子使用AVG()函数返回特定供应商提供的产品的平均价格。 它与上面的SELECT语句相同，但使用了DISTINCT参数，因此平均值只 考虑各个不同的价格:

```mysql
mysql> select avg(distinct prod_price) as avg_price from products where vend_id=1003;
+-----------+
| avg_price |
+-----------+
| 15.998000 |
+-----------+
1 row in set (0.00 sec)
```

## 12.3 组合聚集函数

`SELECT`语句可以根据需要包含多个聚集函数

```mysql
mysql> select count(*) as num_items,min(prod_price) as price_min,max(prod_price) as price_max, avg(prod_price) as price_avg from products;
+-----------+-----------+-----------+-----------+
| num_items | price_min | price_max | price_avg |
+-----------+-----------+-----------+-----------+
|        14 |      2.50 |     55.00 | 16.133571 |
+-----------+-----------+-----------+-----------+
1 row in set (0.00 sec)
```

## 12.4 小结

聚集函数用来汇总数据.MySQL支持一系列聚集函数.它们返回结果一般比在自己的客户机应用程序中计算快得多

# 第13章 分组数据

## 13.1 数组分组

从上一章知道，SQL聚集函数可用来汇总数据。这使我们能够对行进 行计数，计算和与平均数，获得最大和最小值而不用检索所有数据。

但如果要返回每个供应商提供的产品数目怎么办？或者返回只提供 单项产品的供应商所提供的产品，或返回提供10个以上产品的供应商怎 么办？

## 13.2 创建分组

分组是在`SELECT`语句的`GROUP BY`子句中建立的。理解分组最好的办法就是看一个例子。

```mysql
mysql> select vend_id, count(*) as num_prods from products group by vend_id;
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1001 |         3 |
|    1002 |         2 |
|    1003 |         7 |
|    1005 |         2 |
+---------+-----------+
4 rows in set (0.00 sec)
```

分析：上面的SELECT语句指定了两个列，vend_id包含产品供应商的ID， num_prods为计算字段（用COUNT(*)函数建立）。GROUP BY子句指 示MySQL按vend_id排序并分组数据。这导致对每个vend_id而不是整个表 计算num_prods一次。从输出中可以看到，供应商1001有3个产品，供应商 1002有2个产品，供应商1003有7个产品，而供应商1005有2个产品。

在具体使用`GROUP BY`之前，需要知道一些重要的规定。

+ `GROUP BY`子句可以包含任意数目的列。这使得能对分组进行嵌套，为数据分组提供更细致的控制。
+ 如果在`GROUP BY`子句中嵌套了分组，数据将在最后规定的分组上进行汇总。换句话说，在建立分组时，指定的所有列都一起计算 （所以不能从个别的列取回数据）。
+ `GROUP BY` 子句中列出的每个列都必须是检索列或有效的表达式 （但不能是聚集函数）。如果在 SELECT 中使用表达式， 则必须在`GROUP BY`子句中指定相同的表达式。不能使用别名。
+ 除聚集计算语句外，SELECT语句中的每个列都必须在`GROUP BY`子句中给出。
+ 如果分组列中具有NULL值，则NULL将作为一个分组返回。如果列中有多行NULL值，它们将分为一组。
+ `GROUP BY`子句必须出现在WHERE子句之后，`ORDER BY`子句之前。

> 使用ROLLUP 使用WITH ROLLUP关键字，可以得到每个分组以 及每个分组汇总级别（针对每个分组）的值，如下所示：
>
> mysql> mysql> select vend_id, count(*) as num_prods from products group by vend_id with rollup;

## 13.3 过滤分组

除了能用`GROUP BY`分组数据外，MySQL还允许过滤分组，规定包括 哪些分组，排除哪些分组。例如，可能想要列出至少有两个订单的所有 顾客。为得出这种数据，必须基于完整的分组而不是个别的行进行过滤。

MySQL为此目的提供了另外的子 句，那就是`HAVING`子句。`HAVING`非常类似于WHERE。事实上，目前为止所 学过的所有类型的 WHERE 子句都可以用`HAVING`来替代。 唯一的差别是 WHERE过滤行，而`HAVING`过滤分组。

看下面的例子

```mysql
mysql> select cust_id , COUNT(*) as orders from orders group by cust_id having count(*)>=2;
+---------+--------+
| cust_id | orders |
+---------+--------+
|   10001 |      2 |
+---------+--------+
1 row in set (0.00 sec)
```

> HAVINT 和 WHERE的差别 这里有另一种理解方法，WHERE在数据
>
> 分组前进行过滤，HAVING在数据分组后进行过滤。这是一个重
>
> 要的区别，WHERE排除的行不包括在分组中。这可能会改变计
>
> 算值，从而影响HAVING子句中基于这些值过滤掉的分组。

## 13.4 分组和排序

`GROUP BY`和`ORDER BY`的差别

| ORDER BY                                     | GROUP BY                                                 |
| -------------------------------------------- | -------------------------------------------------------- |
| 排序产生的输出                               | 分组行。但输出可能不是分组的顺序                         |
| 任意列都可以使用（甚至非选择的列也可以使用） | 只可能使用选择列或表达式列，而且必须使用每个选择列表达式 |
| 不一定需要                                   | 如果与聚集函数一起使用列（或表达式），则必须使用         |

## 13.5 SELECT语句的顺序

| 子句     | 说明               | 是否必须使用           |
| -------- | ------------------ | ---------------------- |
| SELECT   | 要返回的列或表达式 | 是                     |
| FROM     | 从中检索数据的表   | 仅在从表选择数据时使用 |
| WHERE    | 行级过滤           | 否                     |
| GROUP BY | 分组说明           | 仅在按组计算聚集时使用 |
| HAVING   | 组级过滤           | 否                     |
| ORDER BY | 输出排序顺序       | 否                     |
| LIMIT    | 要检索的行数       | 否                     |

## 13.6 小结

在第12章中，我们学习了如何用SQL聚集函数对数据进行汇总计算。 本章讲授了如何使用GROUP BY子句对数据组进行这些汇总计算，返回每 个组的结果。我们看到了如何使用HAVING子句过滤特定的组，还知道了 `ORDER BY`和`GROUP BY`之间以及`WHERE`和`HAVING`之间的差异。

# 第14章 使用子查询

SQL还允许创建子查询(subquery)，即嵌套在其他查询中的查询。

## 14.2 利用子查询进行过滤

订单存储在两个表中。对于包含订单号、客户ID、 订单日期的每个订单，orders表存储一行。各订单的物品存储在相关的 orderitems表中。orders表不存储客户信息。它只存储客户的ID。实际 的客户信息存储在customers表中。

假如需要列出订购物品TNT2的所有客户，应该怎样检索?下 面列出具体的步骤

1. 检索包含物品TNT2的所有订单的编号。
2. 检索具有前一步骤列出的订单编号的所有客户的ID

可以使用子查询来把3个查询组合成一条语句。

```mysql
mysql> select cust_id from orders where order_num in (select order_num from orderitems where prod_id = 'TNT2');
+---------+
| cust_id |
+---------+
|   10001 |
|   10004 |
+---------+
2 rows in set (0.00 sec)
```

分析:在SELECT语句中，子查询总是从内向外处理。在处理上面的 SELECT语句时，MySQL实际上执行了两个操作。

## 14.3 作为计算字段使用子查询

使用子查询的另一方法是创建计算字段。假如需要显示customers表中每个客户的订单总数。订单与相应的客户ID存储在orders表中。

1. 从customers表中检索客户列表
2. 对于检索出的客户,统计在orders表中的订单数目

```mysql
mysql> select cust_name, cust_state,(select count(*) from orders where orders.cust_id=customers.cust_id) as orders from customers;
+----------------+------------+--------+
| cust_name      | cust_state | orders |
+----------------+------------+--------+
| Coyote Inc.    | MI         |      2 |
| Mouse House    | OH         |      0 |
| Wascals        | IN         |      1 |
| Yosemite Place | AZ         |      1 |
| E Fudd         | IL         |      1 |
+----------------+------------+--------+
5 rows in set (0.00 sec)
```

分析:这条SELECT语句对customers表中每个客户返回3列:

cust_name、cust_state和orders。orders是一个计算字段， 它是由圆括号中的子查询建立的。该子查询对检索出的每个客户执行一 次。在此例子中，该子查询执行了5次，因为检索出了5个客户。

> 相关子查询  涉及外部查询的子查询。

`orders.cust_id = customers_cust_id`

这种类型的子查询称为相关子查询。任何时候只要列名可能有多义 性，就必须使用这种语法(表名和列名由一个句点分隔)

## 14.4 小结

本章学习了什么是子查询以及如何使用它们。子查询最常见的使用 是在WHERE子句的IN操作符中，以及用来填充计算列。

# 第15章 联结表

## 15.1 联结

SQL最强大的功能之一就是能在数据检索查询的执行中联结(join) 表。联结是利用SQL的SELECT能执行的最重要的操作

### 15.1.1 关系表

理解关系表的最好方法是来看一个现实世界中的例子。

假如有一个包含产品目录的数据库表，其中每种类别的物品占一行。 对于每种物品要存储的信息包括产品描述和价格，以及生产该产品的供 应商信息。

现在，假如有由同一供应商生产的多种物品，那么在何处存储供应 商信息(如，供应商名、地址、联系方法等)呢?将这些数据与产品信 息分开存储的理由如下。

+ 因为同一供应商生产的每个产品的供应商信息都是相同的，对每 个产品重复此信息既浪费时间又浪费存储空间。
+ 如果供应商信息改变(例如，供应商搬家或电话号码变动)，只需 改动一次即可。
+ 如果有重复数据(即每种产品都存储供应商信息)，很难保证每次 输入该数据的方式都相同。不一致的数据在报表中很难利用。

关键是，相同数据出现多次决不是一件好事，此因素是关系数据库 设计的基础。关系表的设计就是要保证把信息分解成多个表，一类数据 一个表。各表通过某些常用的值(即关系设计中的关系(relational))互 相关联。

在这个例子中，可建立两个表，一个存储供应商信息，另一个存储 产品信息。vendors表包含所有供应商信息，每个供应商占一行，每个供 应商具有唯一的标识。此标识称为主键(primary key)(在第1章中首次 提到)，可以是供应商ID或任何其他唯一值。

products表只存储产品信息，它除了存储供应商ID(vendors表的主 键)外不存储其他供应商信息。vendors表的主键又叫作products的外键， 它将vendors表与products表关联，利用供应商ID能从vendors表中找出 相应供应商的详细信息。

> 外键  外键为某个表中的一列，它包含另一个表 的主键值，定义了两个表之间的关系。

+ 供应商信息不重复,不浪费时间和空间
+ 如果供应商信息变动,可以只更新vendors表单个记录
+ 数据无重复,显然数据是一致的

> 可伸缩性 能够适应不断增加的工作量而不失败。设 计良好的数据库或应用程序称之为可伸缩性好(scale well)。

### 15.1.2 为什么要使用联结

如果数据存储在多个表中，怎样用单条SELECT语句检索出数据?

答案是使用联结。简单地说，联结是一种机制，用来在一条SELECT 语句中关联表，因此称之为联结。使用特殊的语法，可以联结多个表返 回一组输出，联结在运行时关联表中正确的行。

## 15.2 创建联结

联结的创建非常简单，规定要联结的所有表以及它们如何关联即可。 请看下面的例子:

```mysql
mysql> select vend_name, prod_name, prod_price from vendors, products where vendors.vend_id=products.vend_id;
+-------------+----------------+------------+
| vend_name   | prod_name      | prod_price |
+-------------+----------------+------------+
| Anvils R Us | .5 ton anvil   |       5.99 |
| Anvils R Us | 1 ton anvil    |       9.99 |
| Anvils R Us | 2 ton anvil    |      14.99 |
| LT Supplies | Fuses          |       3.42 |
| LT Supplies | Oil can        |       8.99 |
| ACME        | Detonator      |      13.00 |
| ACME        | Bird seed      |      10.00 |
| ACME        | Carrots        |       2.50 |
| ACME        | Safe           |      50.00 |
| ACME        | Sling          |       4.49 |
| ACME        | TNT (1 stick)  |       2.50 |
| ACME        | TNT (5 sticks) |      10.00 |
| Jet Set     | JetPack 1000   |      35.00 |
| Jet Set     | JetPack 2000   |      55.00 |
+-------------+----------------+------------+
14 rows in set (0.00 sec)
```

分析:我们来考察一下此代码。SELECT语句与前面所有语句一样指定要检索的列。这里，最大的差别是所指定的两个列(prod_name 和prod_price)在一个表中，而另一个列(vend_name)在另一个表中。

可以看到要匹配的两个列以vendors.vend_id和products. vend_id指定。这里需要这种完全限定列名，因为如果只给出vend_id， 则MySQL不知道指的是哪一个(它们有两个，每个表中一个)。

> 完全限定列名 在引用的列可能出现二义性时，必须使用完 全限定列名(用一个点分隔的表名和列名)。如果引用一个 没有用表名限制的具有二义性的列名，MySQL将返回错误。

### 15.2.1 WHERE子句的重要性

利用WHERE子句建立联结关系似乎有点奇怪，但实际上，有一个很充 分的理由。请记住，在一条SELECT语句中联结几个表时，相应的关系是 在运行中构造的。在数据库表的定义中不存在能指示MySQL如何对表进 行联结的东西。你必须自己做这件事情。在联结两个表时，你实际上做 的是将第一个表中的每一行与第二个表中的每一行配对。WHERE子句作为 过滤条件，它只包含那些匹配给定条件(这里是联结条件)的行。

> 笛卡尔积 由没有联结条件的表关系返回的结果为笛卡儿积。检索出的行的数目将是第一个表中的行数乘 以第二个表中的行数。

```mysql
mysql> select vend_name, prod_name, prod_price from vendors,products;
+----------------+----------------+------------+
| vend_name      | prod_name      | prod_price |
+----------------+----------------+------------+
| Anvils R Us    | .5 ton anvil   |       5.99 |
| LT Supplies    | .5 ton anvil   |       5.99 |
| ACME           | .5 ton anvil   |       5.99 |
| Furball Inc.   | .5 ton anvil   |       5.99 |
```

一共84行

### 15.2.2 内部联结

目前为止所用的联结称为**等值联结(equijoin)**，它基于两个表之间的 相等测试。这种联结也称为内部联结。其实，对于这种联结可以使用稍 微不同的语法来明确指定联结的类型。下面的SELECT语句返回与前面例 子完全相同的数据:

```mysql
mysql> select vend_name,prod_name,prod_price from vendors inner join products on vendors.vend_id = products.vend_id;
+-------------+----------------+------------+
| vend_name   | prod_name      | prod_price |
+-------------+----------------+------------+
| Anvils R Us | .5 ton anvil   |       5.99 |
| Anvils R Us | 1 ton anvil    |       9.99 |
| Anvils R Us | 2 ton anvil    |      14.99 |
| LT Supplies | Fuses          |       3.42 |
| LT Supplies | Oil can        |       8.99 |
| ACME        | Detonator      |      13.00 |
| ACME        | Bird seed      |      10.00 |
| ACME        | Carrots        |       2.50 |
| ACME        | Safe           |      50.00 |
| ACME        | Sling          |       4.49 |
| ACME        | TNT (1 stick)  |       2.50 |
| ACME        | TNT (5 sticks) |      10.00 |
| Jet Set     | JetPack 1000   |      35.00 |
| Jet Set     | JetPack 2000   |      55.00 |
+-------------+----------------+------------+
14 rows in set (0.00 sec)
```

分析:此语句中的SELECT与前面的SELECT语句相同，但FROM子句不同。这里，两个表之间的关系是FROM子句的组成部分，以INNER JOIN指定。在使用这种语法时，联结条件用特定的`ON`子句而不是WHERE 子句给出。传递给ON的实际条件与传递给WHERE的相同。

> 使用哪种语法 ANSI SQL规范首选INNER JOIN语法。此外， 尽管使用WHERE子句定义联结的确比较简单，但是使用明确的 联结语法能够确保不会忘记联结条件，有时候这样做也能影响 性能。

### 15.2.3 联结多个表

SQL对一条SELECT语句中可以联结的表的数目没有限制。首先列出所有表，然后定义表之间的关系。

```mysql
mysql> select prod_name,vend_name,prod_price ,quantity from orderitems,products,vendors where products.vend_id = vendors.vend_id and orderitems.prod_id = products.prod_id and order_num = 20005;
+----------------+-------------+------------+----------+
| prod_name      | vend_name   | prod_price | quantity |
+----------------+-------------+------------+----------+
| .5 ton anvil   | Anvils R Us |       5.99 |       10 |
| 1 ton anvil    | Anvils R Us |       9.99 |        3 |
| TNT (5 sticks) | ACME        |      10.00 |        5 |
| Bird seed      | ACME        |      10.00 |        1 |
+----------------+-------------+------------+----------+
4 rows in set (0.00 sec)
```

分析:此例子显示编号为20005的订单中的物品。订单物品存储在orderitems表中。每个产品按其产品ID存储，它引用products 表中的产品。这些产品通过供应商ID联结到vendors表中相应的供应商， 供应商ID存储在每个产品的记录中。这里的FROM子句列出了3个表，而 WHERE子句定义了这两个联结条件，而第三个联结条件用来过滤出订单 20005中的物品。

另一种写法:

```mysql
select prod_name, vend_name, prod_price, quantity from orderitems inner join products on orderitems.prod_id = products.prod_id inner join vendors on products.vend_id = vendors.vend_id and order_num=20005;
```

## 15.3 小结

联结是SQL中最重要最强大的特性，有效地使用联结需要对关系数据 库设计有基本的了解。本章随着对联结的介绍讲述了关系数据库设计的 一些基本知识，包括等值联结(也称为内部联结)这种最经常使用的联 结形式。下一章将介绍如何创建其他类型的联结。

# 第16章 创建高级联结

## 16.1 使用表别名

SQL允许给表名取别名。

+ 缩短SQL语句
+ 允许在单条`SELECT`多次使用相同的表

请看如下例子：

```mysql
mysql> select cust_name,cust_contact from customers as c,orders as o,orderitems as oi where c.cust_id=o.cust_id and oi.order_num=o.order_num and prod_id= 'TNT2';
+----------------+--------------+
| cust_name      | cust_contact |
+----------------+--------------+
| Coyote Inc.    | Y Lee        |
| Yosemite Place | Y Sam        |
+----------------+--------------+
2 rows in set (0.00 sec)
```

## 16.2 使用不同类型的联结

迄今为止，我们使用的只是**内部联结**或**等值连接（equijoin）**的简单联结。来看看其他的联结。**自联结**、**自然联结**和**外部联结**。

### 16.2.1 自联结

假如你发现某物品（ID是**DTNTR**）存在问题，因此想知道生产该物品的供应商生产的其他物品是否也存在问题。首先找到供应商，然后找到供应商的其他物品。

```mysql
mysql> select prod_id,prod_name from products where vend_id=(select vend_id from products where prod_id = 'DTNTR');
+---------+----------------+
| prod_id | prod_name      |
+---------+----------------+
| DTNTR   | Detonator      |
| FB      | Bird seed      |
| FC      | Carrots        |
| SAFE    | Safe           |
| SLING   | Sling          |
| TNT1    | TNT (1 stick)  |
| TNT2    | TNT (5 sticks) |
+---------+----------------+
7 rows in set (0.00 sec)
```

分析:这是一种解决方案，使用了子查询。



下面来看看使用联结的相同查询:

```mysql
mysql> select p1.prod_id,p1.prod_name from products as p1,products as p2 where p1.vend_id=p2.vend_id and p2.prod_id = 'DTNTR';
+---------+----------------+
| prod_id | prod_name      |
+---------+----------------+
| DTNTR   | Detonator      |
| FB      | Bird seed      |
| FC      | Carrots        |
| SAFE    | Safe           |
| SLING   | Sling          |
| TNT1    | TNT (1 stick)  |
| TNT2    | TNT (5 sticks) |
+---------+----------------+
7 rows in set (0.00 sec)
```

分析:此查询中需要的两个实际是相同的表，因此products在`FROM`子句中出现了两次。虽然完全合法，但是具有二义性。

为了解决这个问题，使用了表别名。第一个别名是p1，第二个是p2。

> 使用自联结而不用子查询  自联结通常比子查询快

### 16.2.2 自然联结

**自然联结**排除多次出现，每个列只返回一次

自然联结是这样一种联结，其中你只能选择唯一的列。对其他的表的列使用明确的子集

下面是一个例子

```mysql
mysql> select c.*,o.order_num,o.order_date,oi.prod_id,oi.quantity,oi.item_price from customers as c,orders as o, orderitems as oi where c.cust_id=o.cust_id and oi.order_num=o.order_num and prod_id = 'FB';
```

### 16.2.3 外部联结

许多联结将一个表中的行与另一个表中的行关联。但有时候需要包含没有关联行的那些行。例如，可能需要使用联结来完成以下工作:

+ 对每个客户下了多少订单数进行计数，包括那些至今尚未下订单的客户。
+ 列出所有产品以及订购数量，包括没有人订购的产品

上面例子中，联结包含了那些在相关表中没有关联行的行。这种联结称为**外部联结**。

下面的`SELECT`语句给出了一个外部联结。

```mysql
mysql> select customers.cust_id,orders.order_num from customers left outer join orders on customers.cust_id = orders.cust_id;
+---------+-----------+
| cust_id | order_num |
+---------+-----------+
|   10001 |     20005 |
|   10001 |     20009 |
|   10002 |      NULL |
|   10003 |     20006 |
|   10004 |     20007 |
|   10005 |     20008 |
+---------+-----------+
6 rows in set (0.00 sec)
```

分析:`OUTER JOIN`. 外部联结包含了没有关联的行的行,使用`OUTER JOIN`时候，必须使用`RIGHT`和`LEFT`关键字指定包括所有行的表。

## 16.3 使用带聚集函数的联结

看下面一个例子

```mysql
mysql> select customers.cust_name ,customers.cust_id, count(orders.order_num) as num_ord from customers inner join orders on customers.cust_id = orders.cust_id group by customers.cust_id;
+----------------+---------+---------+
| cust_name      | cust_id | num_ord |
+----------------+---------+---------+
| Coyote Inc.    |   10001 |       2 |
| Wascals        |   10003 |       1 |
| Yosemite Place |   10004 |       1 |
| E Fudd         |   10005 |       1 |
+----------------+---------+---------+
4 rows in set (0.00 sec)
```

## 16.4 使用联结和联结条件

+ 一般我们使用内部联结，但外部联结也是有效的
+ 保证使用正确的联结条件
+ 应该提供联结条件，不然得出笛卡尔积
+ 一个联结中可以包含多个表。

## 16.5 小结

本章是上一章的继续。讲述使用别名，讨论不同的联结类型以及每种联结的语法形式。

# 第17章 组合查询

本章讲述如何利用`UNION`操作符将多条`SELECT`语句组合成一个结果

## 17.1 组合查询

多数SQL查询都只包含一个从或多个表中返回数据的单条`SELECT`语句。MySQL也允许执行多个查询（多个SELECT）.然后将结果作为单个查询结果返回。这些查询结果称为**并（union）**或者**复合查询**

有两种情况，需要使用组合查询：

+ 单个查询中从不同表返回类似结构的数据
+ 对单个表多个查询，按照单个查询返回数据。

> 组合查询和多个WHERE条件 多数情况下，组合相同表的两个 查询完成的工作与具有多个WHERE子句条件的单条查询完成的 工作相同。换句话说，任何具有多个WHERE子句的SELECT语句 都可以作为一个组合查询给出，在以下段落中可以看到这一点。 这两种技术在不同的查询中性能也不同。因此，应该试一下这 两种技术，以确定对特定的查询哪一种性能更好。

## 17.2 创建组合查询

可用 `UNION` 操作符来组合数条SQL查询。

### 17.2.1 使用UNION

`UNION`使用很简单。所需做的只是给出每条`SELECT`语句，在各条语 句之间放上关键字`UNION`。

举一个例子，假如需要价格小于等于5的所有物品的一个列表，而且 还想包括供应商1001和1002生产的所有物品（不考虑价格）。当然，可以 利用WHERE子句来完成此工作，不过这次我们将使用UNION。

```mysql
mysql> select vend_id,prod_id,prod_price from products where prod_price <=5 union select vend_id, prod_id, prod_price from products where vend_id in (1001,1002);
+---------+---------+------------+
| vend_id | prod_id | prod_price |
+---------+---------+------------+
|    1003 | FC      |       2.50 |
|    1002 | FU1     |       3.42 |
|    1003 | SLING   |       4.49 |
|    1003 | TNT1    |       2.50 |
|    1001 | ANV01   |       5.99 |
|    1001 | ANV02   |       9.99 |
|    1001 | ANV03   |      14.99 |
|    1002 | OL1     |       8.99 |
+---------+---------+------------+
8 rows in set (0.00 sec)
```

分析:这条语句由前面的两条SELECT语句组成，语句中用UNION关键 字分隔。UNION指示MySQL执行两条SELECT语句，并把输出组 合成单个查询结果集。

### 17.2.2 UNION规则

+ UNION必须由两条`SELECT`语句组成，语句之间使用关键字`UNION`分隔。
+ UNION的每个查询必须包含相同的列、表达式或聚集函数

### 17.2.3 包含或取消重复的行

`UNION`从查询结果集中自动去除了重复的行。这是`UNION`的默认行为

这是`UNION`的默认行为，但是如果需要，可以改变它。事实上，如果 想返回所有匹配行，可使用`UNION ALL`而不是`UNION`。

```mysql
mysql> select vend_id,prod_id,prod_price from products where prod_price <=5 union all select vend_id, prod_id, prod_price from products where vend_id in (1001,1002);
+---------+---------+------------+
| vend_id | prod_id | prod_price |
+---------+---------+------------+
|    1003 | FC      |       2.50 |
|    1002 | FU1     |       3.42 |
|    1003 | SLING   |       4.49 |
|    1003 | TNT1    |       2.50 |
|    1001 | ANV01   |       5.99 |
|    1001 | ANV02   |       9.99 |
|    1001 | ANV03   |      14.99 |
|    1002 | FU1     |       3.42 |
|    1002 | OL1     |       8.99 |
+---------+---------+------------+
9 rows in set (0.00 sec)
```

### 17.2.4 对组合查询结果排序

`SELECT`语句的输出用`ORDER BY`子句排序。在用UNION组合查询时，只 能使用一条ORDER BY子句，它必须出现在最后一条SELECT语句之后。

## 17.3 小结

本章讲授如何用UNION操作符来组合SELECT语句。利用UNION，可把 多条查询的结果作为一条组合查询返回，不管它们的结果中包含还是不 包含重复。使用UNION可极大地简化复杂的WHERE子句，简化从多个表中 检索数据的工作。

# 第18章 全文本搜索

## 18.1 理解全文本搜索

> 并非所有引擎都支持全文本搜索  MySQL 支持几种基本的数据库引擎。并非所有的引擎都支持本书所描 述的全文本搜索。两个最常使用的引擎为MyISAM和InnoDB， 前者支持全文本搜索，而后者不支持。

## 18.2 使用全文本搜索

为了进行全文本搜索，必须索引被搜索的列，而且要随着数据的改 变不断地重新索引。在对表列进行适当设计后，MySQL会自动进行所有 的索引和重新索引。

在索引之后，SELECT可与Match()和Against()一起使用以实际执行 搜索。

### 18.2.1 启用全文本搜索支持

一般在创建表时启用全文本搜索。CREATE TABLE语句（第21章中介 绍）接受FULLTEXT子句，它给出被索引列的一个逗号分隔的列表。

```mysql
CREATE TABLE productnotes
(
  note_id    int           NOT NULL AUTO_INCREMENT,
  prod_id    char(10)      NOT NULL,
  note_date datetime       NOT NULL,
  note_text  text          NULL ,
  PRIMARY KEY(note_id),
  FULLTEXT(note_text)
) ENGINE=MyISAM;
```

这些列中有一个名为 note_text 的列， 为了进行全文本搜索， MySQL根据子句 FULLTEXT(note_text) 的指示对它进行索引。

### 18.2.2 进行全文本搜索

在索引之后，使用两个函数`Match()`和`Against()`执行全文本搜索， 其中`Match()`指定被搜索的列，`Against()`指定要使用的搜索表达式。

下面举一个例子:

```mysql
mysql> select note_text from productnotes where match(note_text) against('rabbit');
+-----------------------------------------------------------------------------------------------------------------------+
| note_text                                                                                                             |
+-----------------------------------------------------------------------------------------------------------------------+
| Customer complaint: rabbit has been able to detect trap, food apparently less effective now.                          |
| Quantity varies, sold by the sack load.
All guaranteed to be bright and orange, and suitable for use as rabbit bait. |
+-----------------------------------------------------------------------------------------------------------------------+
2 rows in set (0.00 sec)
```

分析:此SELECT语句检索单个列note_text。由于WHERE子句，一个全 文本搜索被执行。Match(note_text)指示MySQL针对指定的 列进行搜索，Against('rabbit')指定词rabbit作为搜索文本。由于有 两行包含词rabbit，这两个行被返回。

刚才的搜索可以简单的使用`LIKE`完成

```mysql
mysql> select note_text from productnotes where note_text like '%rabbit%';
```

### 18.2.3 使用查询扩展

查询扩展用来设法放宽所返回的全文本搜索结果的范围。考虑下面 的情况。你想找出所有提到anvils的注释。只有一个注释包含词anvils， 但你还想找出可能与你的搜索有关的所有其他行， 即使它们不包含词anvils。

使用查询扩展的时候。MySQL对数据和索引进行两边扫描来完成搜索:

+ 首先，进行一个基本的全文本搜索，找出与搜索条件匹配的所有 行；
+ MySQL检查这些匹配行并选择所有有用的词
+ 再次进行全文搜索，不仅使用原来的条件，还使用所有有用的词

下面举一个例子

```mysql
mysql> select note_text from productnotes where match(note_text) against('anvils' with query expansion);
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| note_text                                                                                                                                                |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Multiple customer returns, anvils failing to drop fast enough or falling backwards on purchaser. Recommend that customer considers using heavier anvils. |
| Customer complaint:
Sticks not individually wrapped, too easy to mistakenly detonate all at once.
Recommend individual wrapping.                       |
| Customer complaint:
Not heavy enough to generate flying stars around head of victim. If being purchased for dropping, recommend ANV02 or ANV03 instead. |
| Please note that no returns will be accepted if safe opened using explosives.                                                                            |
| Customer complaint: rabbit has been able to detect trap, food apparently less effective now.                                                             |
| Customer complaint:
Circular hole in safe floor can apparently be easily cut with handsaw.                                                              |
| Matches not included, recommend purchase of matches or detonator (item DTNTR).                                                                           |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
7 rows in set (0.00 sec)
```

分析:这次返回了7行。第一行包含词anvils，因此等级最高。第二 行与anvils无关，但因为它包含第一行中的两个词（customer 和recommend），所以也被检索出来。第3行也包含这两个相同的词，但它 们在文本中的位置更靠后且分开得更远，因此也包含这一行，但等级为 第三。第三行确实也没有涉及anvils（按它们的产品名）。

### 18.2.4 布尔文本支持

`MySQL`支持全文搜索的另外一种形式，称为**布尔方式（boolean mode）**

+ 要匹配的词
+ 要排斥的词
+ 排列提示
+ 表达式分组
+ 另外一些内容

> 即使没有FULLTEXT索引也可以使用  布尔方式不同于迄今为 止使用的全文本搜索语法的地方在于，即使没有定义 FULLTEXT索引，也可以使用它。但这是一种非常缓慢的操作 （其性能将随着数据量的增加而降低）。

举一个例子

```mysql
mysql> select note_text from productnotes where match(note_text) against('heavy' in boolean mode);
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| note_text                                                                                                                                                |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Item is extremely heavy. Designed for dropping, not recommended for use with slings, ropes, pulleys, or tightropes.                                      |
| Customer complaint:
Not heavy enough to generate flying stars around head of victim. If being purchased for dropping, recommend ANV02 or ANV03 instead. |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
2 rows in set (0.00 sec)
```

分析:此全文本搜索检索包含词heavy的所有行（有两行）。其中使用 了关键字`IN BOOLEAN MODE`，但实际上没有指定布尔操作符， 因此，其结果与没有指定布尔方式的结果相同。

为了匹配包含heavy但不包含任意以rope开始的词的行，可使用以下 查询：

```mysql
mysql> select note_text from productnotes where match(note_text) against('heavy -rope*' in boolean mode);
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| note_text                                                                                                                                                |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Customer complaint:
Not heavy enough to generate flying stars around head of victim. If being purchased for dropping, recommend ANV02 or ANV03 instead. |
+----------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```



下面给出所有支持的布尔操作符

| 布尔操作符 | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| +          | 包含，词必须存在                                             |
| -          | 排除，词必须不出现                                           |
| >          | 包含，而且增加等级值                                         |
| <          | 包含，且减少等级值                                           |
| ()         | 把词组成子表达式（允许这些子表达式作为一个组被包含、排除、排列等） |
| ~          | 取消一个词的排序值                                           |
| *          | 词尾的通配符                                                 |
| ""         | 定义一个短语（与单个词的列表不一样，它匹配整个短语以便包含或排除这个短语） |

下面举几个例子，说明某些操作符如何使用：

```mysql
mysql> select note_text from productnotes where match(note_text) against('+heavy +rope' in boolean mode);
```

分析：这个搜索匹配包含词rabbit和bait的行

```mysql
mysql> select note_text from productnotes where match(note_text) against('heavy rope' in boolean mode);
```

分析：没有指定操作符，这个搜索匹配rabbit和bait中的至少一个词的行。

```mysql
mysql> select note_text from productnotes where match(note_text) against("heavy rope" in boolean mode);
```

分析：这个搜索匹配短语`rabbit bait`而不是匹配两个词rabbit和bait

```mysql
mysql> select note_text from productnotes where match(note_text) against('>heavy <rope' in boolean mode);
```

分析：匹配rabbit和carrot，增加前者等级，降低后者等级

```mysql
mysql> select note_text from productnotes where match(note_text) against('+safe +(<combination)' in boolean mode);
```

分析：这个搜搜必须匹配到safe和combination。并且降低后者的等级。

### 18.2.5 全文本搜索的使用说明

+ 在索引全文本数据时，短词被忽略且从索引中排除。短词定义为 那些具有3个或3个以下字符的词（如果需要，这个数目可以更改）。
+ MySQL带有一个内建的非用词（stopword）列表，这些词在索引全文本数据时总是被忽略
+ 许多词出现的频率很高，搜 因此，MySQL规定了一条50%规则，如果一个词出现在50%以上 的行中，则将它作为一个非用词忽略。50%规则不用于IN BOOLEAN MODE。
+ 如果表中的行数少于3行，则全文本搜索不返回结果（因为每个词 或者不出现，或者至少出现在50%的行中）。
+ 忽略词中的单引号。例如，don't索引为dont。
+ 不具有词分隔符（包括日语和汉语）的语言不能恰当地返回全文 本搜索结果。
+ 如前所述，仅在MyISAM数据库引擎中支持全文本搜索。

## 18.3 小结

本章介绍了为什么要使用全文本搜索， 以及如何使用MySQL的 `Match()`和`Against()`函数进行全文本搜索。我们还学习了查询扩展（它 能增加找到相关匹配的机会）和如何使用布尔方式进行更细致的查找控 制。

# 第19章 插入数据

本章介绍如何利用SQL的INSERT语句将数据插入表中。

## 19.1 数据插入

毫无疑问，SELECT是最常使用的SQL语句了(这就是为什么前17章 讲的都是它的原因)。但是，还有其他3个经常使用的SQL语句需要学习。 第一个就是INSERT(下一章介绍另外两个)。

顾名思义，INSERT是用来插入(或添加)行到数据库表的。插入可 以用几种方式使用:

+ 插入完整的行
+ 插入行的一部分
+ 插入多行
+ 插入某些查询的结果

## 19.2 插入完整的行

把数据插入表中的最简单的方法是使用基本的INSERT语法，它要求指定表名和被插入到新行中的值。下面举一个例子:

```mysql
mysql> insert into customers values(null, 'Pep E. LaPew', '100 Main Street', 'Los Angeles', 'CA', '90046', 'USA', NULL, NULL);
Query OK, 1 row affected (0.01 sec)
```

分析:此例子插入一个新客户到customers表。存储到每个表列中的

数据在VALUES子句中给出，对每个列必须提供一个值。如果某 个列没有值(如上面的cust_contact和cust_email列)，应该使用NULL 值(假定表允许对该列指定空值)。

第一列cust_id也为NULL。这是因为每次插入一个新行时，该 列由MySQL自动增量。



虽然这种语法很简单，但并不安全，应该尽量避免使用。

编写INSERT语句的更安全(不过更烦琐)的方法如下:

```mysql
insert into customers(cust_name,cust_contact,cust_email,cust_address,cust_city,cust_state,cust_zip,cust_country)
VALUES('Pep E',NULL,NULL,'100 Main Street','Los Angeles','CA','90046','USA');
```

> 总是使用列的列表  一般不要使用没有明确给出列的列表的 INSERT语句。使用列的列表能使SQL代码继续发挥作用，即使 表结构发生了变化。

> 省略列 如果表的定义允许，则可以在INSERT操作中省略某 些列。省略的列必须满足以下某个条件。
>
> 1.该列定义为允许NULL值(无值或空值)。2.在表定义中给出默认值。这表示如果不给出值，将使用默
>
> 认值。

## 19.3 插入多个行

可以使用多条INSERT语句，甚至一次提交它们，每条语句用一个分号结束

或者，只要每条INSERT语句中的列名(和次序)相同，可以如下组 合各语句:

```mysql
insert into customers(cust_name,cust_address,cust_city,cust_state,cust_zip,cust_country)
VALUES('Pep E','100 Main Street','Los Angeles','CA','90046','USA'),
('M. Martian', '42 Galaxy way','New York', 'NY','11213','USA');
```

## 19.4 插入检索的数据

INSERT一般用来给表插入一个指定列值的行。但是，INSERT还存在 另一种形式，可以利用它将一条SELECT语句的结果插入表中。这就是所 谓的INSERT SELECT，顾名思义，它是由一条INSERT语句和一条SELECT 语句组成的。

假如你想从另一表中合并客户列表到你的customers表。不需要每次 读取一行，然后再将它用INSERT插入，可以如下进行:

```mysql
insert into customers(cust_name,cust_address,cust_city)
SELECT cust_name,cust_address,cust_city from newcustomers;
```

分析:这个例子使用INSERT SELECT从newcustomers中将所有数据导入customers。SELECT语句从newcustomers检索出要插入的值，而不 是列出它们。

INSERT SELECT中SELECT语句可包含WHERE子句以过滤插入的数据。

## 19.5 小结

本章介绍如何将行插入到数据库表。我们学习了使用`INSERT`的几种 方法，以及为什么要明确使用列名，学习了如何用`INSERT SELECT`从其他 表中导入行。下一章讲述如何使用`UPDATE`和`DELETE`进一步操纵表数据。

# 第20 章 更新和删除数据

## 20.1 更新数据

为了更新(修改)表中的数据，可使用UPDATE语句。可采用两种方 式使用UPDATE:

+ 更新表中特定行
+ 更新表中所有行

UPDATE语句非常容易使用，甚至可以说是太容易使用了。基本的 UPDATE语句由3部分组成，分别是:

+ 要更新的表
+ 列名和它们的新值
+ 确定要更新行的过滤条件

```mysql
UPDATE customers set cust_email = 'elemer@fudd.com' where cust_id = 10005;
```

UPDATE语句总是以要更新的表的名字开始。在此例子中，要更新的 表的名字为customers。SET命令用来将新值赋给被更新的列。如这里所 示，SET子句设置cust_email列为指定的值.

UPDATE语句以WHERE子句结束，它告诉MySQL更新哪一行。没有 WHERE子句，MySQL将会用这个电子邮件地址更新customers表中所有 行，这不是我们所希望的。

```mysql
UPDATE customers set cust_name = 'The Fudds',cust_email='elmer@fudd.com' where cust_id=10005;
```

在更新多个列时，只需要使用单个SET命令，每个“列=值”对之间 用逗号分隔(最后一列之后不用逗号)。在此例子中，更新客户10005的 cust_name和cust_email列。

为了删除某个列的值，可设置它为NULL(假如表定义允许NULL值)。

## 20.2 删除数据

为了从一个表中删除(去掉)数据，使用DELETE语句。可以两种方 式使用DELETE:

+ 从表中删除特定的行
+ 从表中删除所有行

下面的语句从customers表中删除一行:

```mysql
DELETE from customers where cust_id=10006;
```

## 20.3 更新和删除的指导原则

前一节中使用的UPDATE和DELETE语句全都具有WHERE子句，这样做的 理由很充分。如果省略了WHERE子句，则UPDATE或DELETE将被应用到表中 190 所有的行。换句话说，如果执行UPDATE而不带WHERE子句，则表中每个行 都将用新值更新。类似地，如果执行DELETE语句而不带WHERE子句，表的所有数据都将被删除。

+ 除非确实打算更新和删除每一行，否则绝对不要使用不带WHERE 子句的UPDATE或DELETE语句。
+ 保证每个表都有主键(如果忘记这个内容，请参阅第15章)，尽可能 像WHERE子句那样使用它(可以指定各主键、多个值或值的范围)。
+ 在对UPDATE或DELETE语句使用WHERE子句前，应该先用SELECT进 行测试，保证它过滤的是正确的记录，以防编写的WHERE子句不 正确。
+ 使用强制实施引用完整性的数据库(关于这个内容，请参阅第15 章)，这样MySQL将不允许删除具有与其他表相关联的数据的行。

## 20.4 小结

我们在本章中学习了如何使用`UPDATE`和`DELETE`语句处理表中的数 据。我们学习了这些语句的语法，知道了它们固有的危险性。本章中还 讲解了为什么`WHERE`子句对`UPDATE`和`DELETE`语句很重要，并且给出了应该 遵循的一些指导原则，以保证数据的安全。

# 第21章  创建和操纵表

## 21.1 创建表

一般有两种创建表的方法：

+ 使用具有交互式创建和管理表的工具
+ 表也可以直接用MySQL语句操纵。

### 21.1.1表创建基础

可以使用`SQL`的`CREATE TABLE`创建表，必须给出下列信息：

+ 新表的名字，在关键字`CREATE TABLE`之后给出。
+ 表列的名字和定义，用逗号分割

下面的MySQL语句创建本书所用的customers表

```mysql
CREATE TABLE `customers` (
  `cust_id` int(11) NOT NULL AUTO_INCREMENT,
  `cust_name` char(50) NOT NULL,
  `cust_address` char(50) DEFAULT NULL,
  `cust_city` char(50) DEFAULT NULL,
  `cust_state` char(5) DEFAULT NULL,
  `cust_zip` char(10) DEFAULT NULL,
  `cust_country` char(50) DEFAULT NULL,
  `cust_contact` char(50) DEFAULT NULL,
  `cust_email` char(255) DEFAULT NULL,
  PRIMARY KEY (`cust_id`)
) ENGINE=InnoDB
```

分析：表名紧跟`CREATE TABLE`之后。实际表定义（所有列）在圆括号之中。各列之间使用逗号隔开。表的主键可以在创建表的时候用`PRIMARY KEY`关键字指定。

### 21.1.2 使用NULL值

每个表列是NULL或者是NOT NULL。

### 21.1.3 主键再介绍

表中每个行必须有唯一的主键值。如果使用单个列，值必须唯一。如果使用多个列，则这些列的组合值必须唯一。

为了创建多个列组成的主键，应该以逗号分隔的列表给出各列名。

```mysql
CREATE TABLE `orderitems` (
  `order_num` int(11) NOT NULL,
  `order_item` int(11) NOT NULL,
  `prod_id` char(10) NOT NULL,
  `quantity` int(11) NOT NULL,
  `item_price` decimal(8,2) NOT NULL,
  PRIMARY KEY (`order_num`,`order_item`),
  KEY `fk_orderitems_products` (`prod_id`),
  CONSTRAINT `fk_orderitems_orders` FOREIGN KEY (`order_num`) REFERENCES `orders` (`order_num`),
  CONSTRAINT `fk_orderitems_products` FOREIGN KEY (`prod_id`) REFERENCES `products` (`prod_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```

### 21.1.4 使用AUTO_INCREMENT

`AUTO_INCREMENT`告诉MySQL，本列每当增加一行时候自动增量。每次执行一个`INSERT`的时候，MySQL自动对该列增量。

每个表只允许一个`AUTO_INCREMENT`，而且必须被索引（例如成为主键）

### 21.1.5 指定默认值

如果插入时没有给出值，MySQL允许使用默认值。通过`CREATE TABLE`语句的列定义中的`DEFAULT`关键字指定。 

## 21.2 更新表

为了更新表定义，可以使用`ALTER TABLE`语句。

+ 在`ALTER TABLE`之后给出要更改的表名。
+ 所做更改的列表

下面例子给表添加一个列：

```mysql
alter TABLE vendors add vend_phone char(20);
```

删除刚刚添加的列.

```mysql
alter table vendors drop column vend_phone;
```

## 21.3 删除表

删除表非常简单，使用`DROP TABLE`语句即可。

```mysql
drop table customers2;
```

## 21.4 重命名表

```mysql
RENAME TABLE customers2 TO customers;

RENAME TABLE table1 TO table2, table3 TO table4;
```

## 21.5 小结

本章介绍了几条新SQL语句。

`CREATE TABLE` 用来创建新表

`ALTER TABLE`用来更改表列

`DROP TABLE`用来完整的删除一个表

# 第22章 使用视图

本章介绍视图是什么，怎么工作，如何使用。

## 22.1 视图

视图是虚拟的表。视图只包含使用时动态检索数据的查询。

### 22.1.1 为什么使用视图

+ 重用SQL语句
+ 检查复杂的SQL操作
+ 使用表的组成部分而不是整个表
+ 保护数据
+ 更改数据格式和表示

在视图创建之后，可以用与表基本相同的方式利用它们。可以对视 图执行SELECT操作，过滤和排序数据，将视图联结到其他视图或表，甚 至能添加和更新数据（添加和更新数据存在某些限制。关于这个内容稍 后还要做进一步的介绍）。

重要的是知道视图仅仅是用来查看存储在别处的数据的一种设施。 视图本身不包含数据，因此它们返回的数据是从其他表中检索出来的。 在添加或更改这些表中的数据时，视图将返回改变过的数据。

> 性能问题    因为视图不包含数据，所以每次使用视图时，都必须处理查询执行时所需的任一个检索。如果你用多个联结和过滤创建了复杂的视图或者嵌套了视图，可能会发现性能下降得很厉害。

### 22.1.2 关于视图的规则和限制

+ 与表一样，视图必须唯一命名
+ 对于可以创建的视图数目没有限制
+ 为了创建视图，必须具有足够的访问权限。这些限制通常由数据库管理人员授予
+ 视图可以嵌套，即可以利用从其他视图中检索数据的查询来构造一个视图。
+ ORDER BY可以用在视图中，但如果从该视图检索数据SELECT中也含有ORDER BY，那么该视图中的ORDER BY将被覆盖。
+ 视图不能索引，也不能有关联的触发器或默认值。
+ 视图可以和表一起使用。例如，编写一条联结表和视图的SELECT语句。


## 22.2 使用视图

+ 视图用`CREATE VIEW`语句来创建。
+ 使用`SHOW CREATE VIEW viewname；`来查看创建视图的语句。
+ 用DROP删除视图，其语法为`DROP VIEW viewname;`。
+ 更新视图时，可以先用DROP再用CREATE，也可以直接用CREATE OR REPLACE VIEW。如果要更新的视图不存在，则第2条更新语句会创建一个视图；如果要更新的视图存在， 则第 2条更新语句会替换原有视图。


### 22.2.1 用视图简化复杂的联结

视图的最常见的应用之一是隐藏复杂的SQL，这通常都会涉及联结。 请看下面的例子：

```mysql
create view productcustomers as SELECT cust_name,cust_contact, prod_id from customers,orders,orderitems
where customers.cust_id = orders.cust_id and orderitems.order_num = orders.order_num;
```

分析：这条语句创建一个名为productcustomers的视图，它联结三个 表， 以返回已订购了任意产品的所有客户的列表。 如果执行 SELECT * FROM productcustomers，将列出订购了任意产品的客户。

为了检索订购了产品`TNT2`的客户，可以如下进行

```mysql
SELECT cust_name,cust_contact from productcustomers where prod_id = 'TNT2';
```

### 22.2.5 更新视图

迄今为止的所有视图都是和SELECT语句使用的。然而，视图的数据 能否更新？答案视情况而定。

通常， 视图是可更新的（即， 可以对它们使用 INSERT 、 UPDATE 和 DELETE）。更新一个视图将更新其基表（可以回忆一下，视图本身没有数 据）。如果你对视图增加或删除行，实际上是对其基表增加或删除行。

但是，并非所有视图都是可更新的。基本上可以说，如果MySQL不 能正确地确定被更新的基数据，则不允许更新（包括插入和删除）。这实 际上意味着，如果视图定义中有以下操作，则不能进行视图的更新：

+ 分组（使用了`GROUP BY`和`HAVING`）
+ 联结
+ 子查询
+ 并
+ 聚集函数（`MIN()`/`COUNT()`/`SUM()`等）
+ `DISTINCT`
+ 导出（计算）列

## 22.3 小结

视图为虚拟的表。它们包含的不是数据而是根据需要检索数据的查 询。视图提供了一种MySQL的SELECT语句层次的封装，可用来简化数据 处理以及重新格式化基础数据或保护基础数据。

# 第23章 使用存储过程

本章介绍什么是存储过程，为什么要使用存储过程以及如何使用存 储过程，并且介绍创建和使用存储过程的基本语法。

## 23.1 存储过程

迄今为止，使用的大多数SQL语句都是针对一个或多个表的单条语 句。并非所有操作都这么简单，经常会有一个完整的操作需要多条语句 才能完成。例如，考虑以下的情形。

+ 为了处理订单，需要核对以保证库存中有相应的物品。
+ 如果库存有物品，这些物品需要预定以便不将它们再卖给别的人，并且要减少可用的物品数量以反映正确的库存量。
+ 库存中没有的物品需要订购，这需要与供应商进行某种交互。
+ 关于哪些物品入库（并且可以立即发货）和哪些物品退订，需要通知相应的客户。


可以创建存储过程。存储过程简单来说，就是为以后的使用而保存 的一条或多条MySQL语句的集合。可将其视为批文件，虽然它们的作用 不仅限于批处理。

## 23.2 为什么要使用存储过程

既然我们知道了什么是存储过程，那么为什么要使用它们呢？有许 多理由，下面列出一些主要的理由。

+ 通过把处理封装在容易使用的单元中，简化复杂的操作（正如前 面例子所述）。
+ 由于不要求反复建立一系列处理步骤，这保证了数据的完整性。 如果所有开发人员和应用程序都使用同一（试验和测试）存储过 程，则所使用的代码都是相同的。
+ 简化对变动的管理。如果表名、列名或业务逻辑（或别的内容） 有变化，只需要更改存储过程的代码。使用它的人员甚至不需要 知道这些变化。
+ 提高性能。因为使用存储过程比使用单独的SQL语句要快。
+ 存在一些只能用在单个请求中的MySQL元素和特性，存储过程可以使用它们来编写功能更强更灵活的代码

## 23.3 使用存储过程

### 23.3.1 执行存储过程

MySQL称存储过程的执行为调用，因此MySQL执行存储过程的语句 为CALL。CALL接受存储过程的名字以及需要传递给它的任意参数。请看 以下例子：

```mysql
call productpricing(@pricelow,@pricehigh,@priceaverage);
```

### 23.3.2 创建存储过程

正如所述，编写存储过程并不是微不足道的事情。为让你了解这个 过程，请看一个例子——一个返回产品平均价格的存储过程。

```mysql
create PROCEDURE productpricing()
BEGIN
	SELECT AVG(prod_price) as item_price
	from products;
END;
```

分析：我们稍后介绍第一条和最后一条语句。 此存储过程名为 productpricing ， 用 CREATE PROCEDURE productpricing() 语 句定义。如果存储过程接受参数，它们将在()中列举出来。此存储过程没 有参数，但后跟的()仍然需要。BEGIN和END语句用来限定存储过程体，过 程体本身仅是一个简单的SELECT语句（使用第12章介绍的Avg()函数）。

如何使用这个存储过程？

```mysql
call productpricing();
```

分析：CALL productpricing(); 执行刚创建的存储过程并显示返回的结果。因为存储过程实际上是一种函数，所以存储过程名后 需要有()符号（即使不传递参数也需要）。

### 23.3.3 删除存储过程

```mysql
drop PROCEDURE productpricing;
```

### 23.3.4 使用参数

productpricing只是一个简单的存储过程，它简单地显示SELECT语 句的结果。一般，存储过程并不显示结果，而是把结果返回给你指定的变量。

> 变量 内存中一个特定的位置，用来临时存储数据。

以下是productpricing的修改版本（如果不先删除此存储过程，则 不能再次创建它）：

```mysql
create PROCEDURE productpricing(
	OUT pl DECIMAL(8,2),
	OUT ph DECIMAL(8,2),
	OUT pa DECIMAL(8,2)
)
BEGIN
	SELECT MIN(prod_price)
	into pl
	from products;
	SELECT MAX(prod_price)
	into ph
	from products;
	SELECT AVG(prod_price)
	into pa
	from products;
END;
```

分析：此存储过程接受3个参数：pl存储产品最低价格，ph存储产品 最高价格，pa存储产品平均价格。每个参数必须具有指定的类 型，这里使用十进制值。关键字OUT指出相应的参数用来从存储过程传出 一个值（返回给调用者）。MySQL支持IN（传递给存储过程）、OUT（从存 储过程传出，如这里所用）和INOUT（对存储过程传入和传出）类型的参 数。存储过程的代码位于BEGIN和END语句内，如前所见，它们是一系列 SELECT语句，用来检索值，然后保存到相应的变量（通过指定INTO关键 字）。

为了调用这个存储过程。必须指定3个变量名。

```mysql
CALL productpricing(@pricelow, @pricehigh, @priceaverage);
```

为了获得三个值，可以使用以下语句：

```mysql
SELECT @pricelow, @pricehigh, @priceaverage;
```

### 23.3.5 建立只能存储过程

迄今为止使用的所有存储过程基本上都是封装MySQL简单的SELECT 语句。虽然它们全都是有效的存储过程例子，但它们所能完成的工作你 直接用这些被封装的语句就能完成（如果说它们还能带来更多的东西，那就是使事情更复杂）。只有在存储过程内包含业务规则和智能处理时， 它们的威力才真正显现出来。

考虑这个场景。你需要获得与以前一样的订单合计，但需要对合计 增加营业税，不过只针对某些顾客（或许是你所在州中那些顾客）。那么， 你需要做下面几件事情：

+ 获得合计
+ 把营业税有条件的添加到合计
+ 返回合计

```mysql
create PROCEDURE ordertotal(
	IN onumber INT,
	OUT ototal DECIMAL(8,2)
)
BEGIN
	DECLARE total DECIMAL(8.2);
	DECLARE taxrate INT DEFAULT 6;
	SELECT SUM(item_price*quantity)
	FROM orderitems
	WHERE order_num=onumber
	into total;
	
	SELECT total into ototal;
END
```

分析：在存储过程复杂性增加时，这样做特别重要。添加了另外一个 参数taxable，它是一个布尔值（如果要增加税则为真，否则为假）。在 存储过程体中，用DECLARE语句定义了两个局部变量。DECLARE要求指定 变量名和数据类型，它也支持可选的默认值（这个例子中的taxrate的默 认被设置为6%）。SELECT语句已经改变，因此其结果存储到total（局部 变量）而不是ototal。IF语句检查taxable是否为真，如果为真，则用另 一SELECT语句增加营业税到局部变量total。最后，用另一SELECT语句将 total（它增加或许不增加营业税）保存到ototal。

看看调用

```mysql
call ordertotal(20005 ,@total);
select @total;
```

### 23.3.6 检查存储过程

为显示用来创建一个存储过程的 CREATE 语句， 使用 `SHOW CREATE PROCEDURE`语句：

为了获得包括何时、由谁创建等详细信息的存储过程列表，使用`SHOW PROCEDURE STATUS`。

> 限制过程状态结果  SHOW PROCEDURE STATUS列出所有存储过程。为限制其输出，例如：show PROCEDURE STATUS where Db='crashcourse';
>

## 23.4 小结

本章介绍了什么是存储过程以及为什么要使用存储过程。我们介绍 了存储过程的执行和创建的语法以及使用存储过程的一些方法。

# 第24章 使用游标

## 24.1 游标

有时，需要在检索出来的行中前进或后退一行或多行。这就是使用 游标的原因。

游标（cursor）是一个存储在MySQL服 它不是一条SELECT语句，而是被该语句检索出来的结果集。

## 24.2 使用游标

+ 在能够使用游标前，必须声明（定义）它。这个过程实际上没有 检索数据，它只是定义要使用的SELECT语句。

+ 一旦声明后， 必须打开游标以供使用。 这个过程用前面定义的SELECT语句把数据实际检索出来。

+ 对于填有数据的游标，根据需要取出（检索）各行。

+ 在结束游标使用时，必须关闭游标。

在声明游标后，可根据需要频繁地打开和关闭游标。在游标打开后， 可根据需要频繁地执行取操作。

### 24.2.1 创建游标

游标用DECLARE语句创建（参见第23章）。DECLARE命名游标，并定义 相应的SELECT语句，根据需要带WHERE和其他子句。例如，下面的语句定 义了名为ordernumbers的游标，使用了可以检索所有订单的SELECT语句。

```mysql
create PROCEDURE processorders()
BEGIN
	DECLARE ordernumbers CURSOR
	FOR SELECT order_num from orders;
END;
```

分析：这个存储过程并没有做很多事情，DECLARE语句用来定义和命 名游标，这里为ordernumbers。存储过程处理完成后，游标就 消失（因为它局限于存储过程）。

### 24.2.2 打开和关闭游标

游标用`OPEN CURSOR`语句来打开：

```mysql
open ordernumbers;
```

分析：在处理OPEN语句时执行查询，存储检索出的数据以供浏览和滚 动。

游标处理完成后，应当使用如下语句关闭游标：

```mysql
CLOSE ordernumbers;
```

分析：CLOSE释放游标使用的所有内部内存和资源，因此在每个游标 不再需要时都应该关闭。

在一个游标关闭后，如果没有重新打开，则不能使用它。但是，使用声明过的游标不需要再次声明，用OPEN语句打开它就可以了。

下面是前面例子的修改版本：

```mysql
create PROCEDURE processorders()
BEGIN
	DECLARE ordernumbers CURSOR
	FOR SELECT order_num from orders;
	
	open ordernumbers;
	CLOSE ordernumbers;
END;
```

分析：这个存储过程声明、打开和关闭一个游标。但对检索出的数据 什么也没做。

### 24.2.3 使用游标数据

在一个游标被打开后，可以使用`FETCH`语句分别访问它的每一行。 FETCH指定检索什么数据（所需的列），检索出来的数据存储在什么地方。 它还向前移动游标中的内部行指针，使下一条FETCH语句检索下一行（不 重复读取同一行）。

```mysql
create PROCEDURE processorders()
BEGIN
	DECLARE o INT;

	DECLARE ordernumbers CURSOR
	FOR SELECT order_num from orders;
	
	open ordernumbers;
	
	FETCH ordernumbers into o;
	
	CLOSE ordernumbers;
END;
```

分析：其中FETCH用来检索当前行的order_num列（将自动从第一行开 始）到一个名为o的局部声明的变量中。对检索出的数据不做 任何处理。

下一个例子中，循环检索数据，从第一行到最后一行。

```mysql
create PROCEDURE processorders()
BEGIN
	DECLARE o INT;
	DECLARE done TINYINT default 0;

	DECLARE ordernumbers CURSOR
	
	FOR SELECT order_num from orders;
	DECLARE CONTINUE HANDLER for SQLSTATE '02000' SET done=1;
	
	open ordernumbers;
	repeat
	
		FETCH ordernumbers into o;
	UNTIL done=1 END REPEAT;
	
	CLOSE ordernumbers;
END;
```

分析：与前一个例子一样，这个例子使用 FETCH 检索当前 order_num 到声明的名为o的变量中。但与前一个例子不一样的是，这个 例子中的FETCH是在REPEAT内，因此它反复执行直到done为真（由UNTIL done END REPEAT;规定）。为使它起作用，用一个DEFAULT 0（假，不结 束）定义变量done。

## 24.3 小结

本章介绍了什么是游标以及为什么要使用游标，举了演示基本游标 使用的例子，并且讲解了对游标结果进行循环以及逐行处理的技术。
