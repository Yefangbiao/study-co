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
