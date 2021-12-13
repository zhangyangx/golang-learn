package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-learn/src/7_db/mysql/global"
)

func main() {
	// 初始化数据库
	global.InitDB()

	person := SelectOne(2)
	fmt.Println(person)

	personList := SelectList(1)
	fmt.Println(personList)

	preparePersonList := Prepare(1)
	fmt.Println(preparePersonList)

}

// SelectOne 查询单行记录
func SelectOne(id int) global.Person {
	r := global.DB.QueryRow("select user_id, username, sex, email from person where user_id=?", id)
	var person global.Person
	// 将查询结果映射到struct
	err := r.Scan(&person.UserId, &person.Username, &person.Sex, &person.Email)
	if err != nil {
		fmt.Println("数据映射失败")
	}
	return person
}

// SelectList 查询多行记录
func SelectList(id int) []global.Person {
	rows, err := global.DB.Query("select user_id, username, sex, email from person where user_id > ?", id)
	if err != nil {
		fmt.Println("查询失败")
	}
	var result []global.Person
	for rows.Next() {
		var person global.Person
		err := rows.Scan(&person.UserId, &person.Username, &person.Sex, &person.Email)
		if err != nil {
			fmt.Println("数据映射失败")
		}
		result = append(result, person)
	}
	return result
}

// Prepare 方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
func Prepare(id int) []global.Person {
	// 1 写sql语句
	sqlStr := "select user_id, username, sex, email from person where user_id > ?"
	// 2.发送给mysql服务器进行准备，返回一个状态
	stem, err := global.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
	}
	defer func(stem *sql.Stmt) {
		err := stem.Close()
		if err != nil {
			fmt.Printf("close failed, err:%v\n", err)
		}
	}(stem)

	rows, err := stem.Query(id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("close failed, err:%v\n", err)
		}
	}(rows)

	var result []global.Person
	// 循环读取得到的结果
	for rows.Next() {
		var person global.Person
		err := rows.Scan(&person.UserId, &person.Username, &person.Sex, &person.Email)
		if err != nil {
			fmt.Println("数据映射失败")
		}
		result = append(result, person)
	}

	return result
}
