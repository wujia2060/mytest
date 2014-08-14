package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func testSelect() {
	db, err := sql.Open("mysql", "metabus:metabus@tcp(127.0.0.1:3306)/metabus?charset=utf8")
	if err != nil {
		fmt.Printf("connect err")
	}
	rows, err1 := db.Query("select num,name from test")
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	defer rows.Close()
	fmt.Println("")
	cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println("")
	var num int
	var name string
	for rows.Next() {
		if err := rows.Scan(&num, &name); err == nil {
			fmt.Print(num)
			fmt.Print("\t")
			fmt.Print(name)
			fmt.Print("\t\r\n")
		}
	}
}
func main() {
	testSelect()
}
