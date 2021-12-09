package main

import (
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
