package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-learn/src/7_db/mysql/global"
)

func main() {
	// 初始化数据库
	global.InitDB()
	Delete(2)
}

func Delete(id int) {
	res, err := global.DB.Exec("delete from person where user_id=?", id)
	if err != nil {
		fmt.Println("删除失败")
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("delete succ:", row)
}
