package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-learn/src/7_db/mysql/global"
)

func main() {
	global.InitDB()
	r, err := global.DB.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "man", "stu02@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ:", id)
}
