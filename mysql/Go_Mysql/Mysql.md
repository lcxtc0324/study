## Mysql

#### DCL

##### 创建用户

```mysql
create user golang@'%' identified by 'redhat';
```

##### 赋予权限

```mysql
grant all privileges on golang.* to golang@'%' identified by 'redhat';
```

#### DDL

##### 创建表

```mysql
create table 表名(列名 列类型[列的修饰]) engine=innodb default charset utf8mb4;

create table user (
  id bigint primary key auto_increment,
  name varchar(32) not null default '',
  password varchar(1024) not null default '',
  sex boolean not null default true,
  birthday date not null,
  addr text not null default '',
  tel varchar(32) not null default '',
  index idx_name(name),
  index idx_birthday(birthday)
 )engine=innodb default charset utf8mb4;
```

<img src="/Users/luo.cx/Desktop/Go_Mysql/Image/截屏2020-12-07 下午12.59.24.png" alt="截屏2020-12-07 下午12.59.24" style="zoom: 67%;" />



清空表/重建表: truncate table 表名; // 删除数据 DDL

库名，表名，列名: 小写英文字母， _

分表分库: log_数字/日期

##### mysql数据类型与Go语言的对应表

id  整数类型:
int => int32, bigint => int 64

------

name 字符串类型 ：

string []byte

char(length) 需定义长度 最大10

varchar(length) 需定义长度 //*常用

text  65535 byte //*常用

​	longtext 4G

blob

------

password

sex bool

​	true  男

​	false 女

------

birthday

time.time

datetime

------

addr

​	text

------

tel

​	varchar(32)

##### 查看详情

```mysql
desc 表名;
```

##### 删除表

```mysql
drop table 表名;
```

##### 常用数据类型:

数值类型

​	布尔类型	boolean *// 常用

​	整形	int (4byte) , bigint (8byte) *// 常用

​	浮点型

​		float *// 常用

​		double

​		decimal(m, d) 	 16 m: 有效位数 小数点之前和小数点之后的数字的个数最大值										   d: 小数点后最大位数



字符类型

​	char

​	varchar *// 常用

​	text  65535 (tinytext 255K, mediumtext 16M, longtext 4G) *// 常用

​	blob

​	json => 关系型数据库支持文档格式

​	enum



日期类型

​	date 年月日

​	time 时分秒

​	datetime 年月日-时分秒  *// 常用

​	timestamp 年月日时分秒 (更新时会自动更新为当前时间)

​	

##### 修饰

​	主键: PRIMARY KEY

​	主键自动增长:  int/bigint 类型 AUTO_INCREMENT

​	唯一: UNIQUE KEY 	

​	非NULL: NOT NULL 

​		默认值: DEFAULT value

​					  DEFAULT 0

  					DEFAULT ' '

​		注释: COMMENT

------

##### 修改表

列 重命名, 添加, 类型修改, 删除列,

alter table 表名 add/change/drop column 列名  列类型  列修饰

```
alter table user add column created_at datetime not null;
```

------

##### 索引

​	针对查询字段,不会针对枚举类型创建

​		用户 status int 1: 未激活 2: 正常使用 3: 锁定 4: 删除

```
create index 索引名称 [using] 索引类型 (创建的列)on 表名;
create unique index 唯一索引名称 [using 索引类型](创建的列) on 表名;
联合索引(多个KEY)
hash * btree rtree
create index idx_name on user(name);
drop index 索引名称 on 表名;

```

------

#### DQL

```mysql
	select * from 表名;
	select * from WHERE 条件；
			条件: 列名 基准对象 比较
				比较:
					关系运算
						>< = != >= <=
					布尔运算
						and or not
					like 
						以某个字符串开头 like 'substr%'
						以某个字符串结尾 like '%substr'
						包含字符串 like '%substr%'
						
						% 0个或任意个字符
						_ 一个或者固定数量的任意字符 substr 
						
					like binary
						name kk, AK
						
					in (v1, v2, vn) 列表
					字符串类型
						关系运算
						包含内容 like
					数值类型
						关系运算
						
						函数
						四则运算 + - * / %
					时间类型
						关系运算
					bool类型
```

------

#### DML

​	添加

​		insert into 表名(列名) values(值);

​		insert into 表名 values(值); // 所有列都添加数据

​		NOT NULL 但是设置为 DEFAULT 该情况插入数据必须指定该列	

```
insert into user values(1, 'luocx','123','true','1999-10-11','北京','13943332212','2020-12-0inser7')
```

​	修改

​		update 表名

​		SET 列名=值, 列名=；

```
update user set addr='北京' where addr ='';
update user set addr = '北京' where addr='';
```

​	删除

​		delete from 表名 [where 条件]

```
delete from user where tel = '';
```

​		truncate table 表名;  // 删除数据

​	查询

```mysql
select * from user where tel='15943330202';
select * from user where name like '%kk%'; 包含kk
select * from user where birthday >= '2011-01-01 00:00:00';
select * from user where birthday >= '2011-01-01 00:00:00' and name like 'kk%';
select * from user where addr = '北京' or addr = '西安';
select * from user where addr in('西安', '北京');
select * from user where  addr = ''; #查询空值
# 限制查询结果条数
select * from user limit 3 ;  #限制查询 , 用于分页
select * from user offset 3 ; #
limit N offset M ;

```

##### 分页:

​	pageSize 每页显示数量
​	pageNum 显示第几页 1,2,3,....
​	limit pageSize offset pageNum (pageNum -1)
​	
select * from 表名 where 条件 limit N offset M;

------

##### 排序:

order by colname [asc/desc],
selct * from user oder by birthday asc; #升序 （默认）
selct * from user oder by birthday desc; #降序
selct * from user oder by birthday asc, name desc; #多列
where 条件 order by limit N  offset M;

------

##### 聚合查询:

groupby

select addr, count(*) from user group by addr;

count(*)

sum(column)

avg(column)

max(column)

min(column)

```mysql
select min(birthday), max(birthday), count(*) from user;
select count(*) from user  where 
select  addr, name, count(*) from user group by addr;
```

------

#### 多表操作

##### 	join 多个表之间存在关联关系(查询)

​	两张表都有相同的关联才能查询

```mysql
select * from password join shadow on password.id=shadow;
```

![](/Users/luo.cx/Desktop/Go_Mysql/Image/截屏2020-12-07 下午4.27.00.png)

------

##### 分表:

​	每个月一张表

​	集合 交际, 并集, 差集

​	union

​	accesslog_01

​	accesslog_02

​	

------

## Go操作MySQL

database/sql => 定义对数据库操作接口, 未实现针对数据库操作功能

github.com/go-sql-driver/mysql



#### 操作:

##### 	查询

​		DQL

​		Query

##### 	修改

​		DDL, DCL, DML

​		Exec		 

1. ##### 使用驱动

   a. 选择驱动

   b. 初始化导入驱动

   ![截屏2020-12-11 下午4.33.08](/Users/luo.cx/Desktop/Go_Mysql/Image/截屏2020-12-11 下午4.33.08.png)

2. ##### 打开数据库(连接池)

3. ##### 操作:

   a. 修改

   ​	Exec

   b. 查询

   ​	Query 提取

   ​	QueryRow  单行提取

   ​	for Rows.Next  判断是否有下一行

   ​    Rows.Scan 扫描结果

4. ##### 关闭资源

   a. 查询

   ​	Rows.Close()

5. ##### 进程退出

   关闭数据库(连接池)

   db.Close()

6. 数据库预处理 (防sql注入)

   使用 ? 占位符预处理方式。。

   `通常为where条件开始,表情,表字段 ，表属性，关键字无法替代`

![截屏2020-12-11 下午3.41.23](/Users/luo.cx/Desktop/Go_Mysql/Image/截屏2020-12-11 下午3.41.23.png)

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	driverName := "mysql"
	dsn := "golang:redhat@tcp(10.34.90.20:3306)/golang?charset=utf8mb4&loc=Local&parseTime=true" //datastore name 数据库连接信息, 使用协议，用户&密码，数据库，连接参数 (parseTime=ture 时间相关)
	db, err := sql.Open(driverName, dsn)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()


	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	name := "%kk%" //sql 注入
	sql  := `
			select id, name, password, sex, birthday, addr, tel
			from user
			where name like ?
			order by ? desc
			limit ? offset ?
	`
	//fmt.Println(sql)

	// 操作
	rows, err := db.Query(sql, name, "birthday", 3, 0) //数据库的预处理方式
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id int64
			name string
			password string
			sex bool
			birthday *time.Time
			addr string
			tel string
		)
		err := rows.Scan(&id, &name, &password, &sex, &birthday, &addr, &tel)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(id, name, password, sex,  birthday, addr, tel)
		}
	}
	var id int64
	err = db.QueryRow("select id from user order by id").Scan(&id)
	fmt.Println(err, id)
```

Exec

