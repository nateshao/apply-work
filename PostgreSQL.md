<center><h1>PostgreSQL 练习</h1></center>

[TOC]



## Windows安装

参考：https://www.runoob.com/postgresql/windows-install-postgresql.html

## 遇到错误

参考https://blog.csdn.net/sinat_36226553/article/details/100750378 已经解决

##  PostgreSQL 创建数据库

PostgreSQL 创建数据库可以用以下三种方式：

- 1、使用 **CREATE DATABASE** SQL 语句来创建。
- 2、使用 **createdb** 命令来创建。
- 3、使用 **pgAdmin** 工具。

-  CREATE DATABASE 创建数据库

CREATE DATABASE 命令需要在 PostgreSQL 命令窗口来执行，语法格式如下：

```sql
CREATE DATABASE dbname;
```

例如，我们创建一个 runoobdb 的数据库：

```sql
postgres=# CREATE DATABASE runoobdb;
```

## PostgreSQL 选择数据库

- 使用 **\l** 用于查看已经存在的数据库：

接下来我们可以使用 **\c + 数据库名** 来进入数据库：

```sql
postgres=# \c runoobdb
You are now connected to database "runoobdb" as user "postgres".
runoobdb=# 
```

### pgAdmin 工具

pgAdmin 工具更简单了，直接点击数据库选择就好了，还可以查看一些数据库额外的信息：

![img](https://nateshao-blog.oss-cn-shenzhen.aliyuncs.com/img/029E4150-C4C3-47BE-8CCE-9267335EDBBB.jpg)

## PostgreSQL 删除数据库

PostgreSQL 删除数据库可以用以下三种方式：

- 1、使用 **DROP DATABASE** SQL 语句来删除。
- 2、使用 **dropdb** 命令来删除。
- 3、使用 **pgAdmin** 工具。

**注意：**删除数据库要谨慎操作，一旦删除，所有信息都会消失。

例如，我们删除一个 runoobdb 的数据库：

```sql
postgres=# DROP DATABASE runoobdb;
```



## TODO...

参照 [菜鸟教程](https://www.runoob.com/postgresql/postgresql-tutorial.html)





