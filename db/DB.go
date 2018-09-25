package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:zhangyizheng@/chat")
	if err != nil {
		fmt.Println(nil)
		return
	}

	// result, err := db.Exec("CREATE TABLE person(name varchar(255), age integer )")
	//result, err := db.Exec("DROP TABLE person"）
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(result)

	//result, err := db.Exec("")
}
