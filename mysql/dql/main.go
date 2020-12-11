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
}

