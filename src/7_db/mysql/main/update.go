package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-learn/src/7_db/mysql/global"
)

func main() {

	// 初始化数据库
	global.InitDB()

	person := global.Person{UserId: 2, Username: "张三"}
	Update(&person)

}

func Update(person *global.Person) {
	res, err := global.DB.Exec("update person set username=? where user_id=?", person.Username, person.UserId)
	if err != nil {
		fmt.Println("更新失败")
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("update succ:", row)
}
