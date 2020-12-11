package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	driverName := "mysql"
	dsn := "golang:redhat@tcp(10.34.90.20:3306)/golang?parseTime=true&loc=Local&charset=utf8mb4"
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


	sql := `
		INSERT INTO user(name, password, birthday) VALUES('jjs', 'xxxxxx', '1991-11-11');
	`

	result, err := db.Exec(sql); if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}

	sql := `
	UPDATE 
	`
	result, err := db.Exec()

}
