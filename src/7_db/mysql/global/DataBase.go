package global

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	DB = database
	// 注意这行代码要写在上面err判断的下面
	//defer func(database *sql.DB) {
	//	err := database.Close()
	//	if err != nil {
	//		fmt.Println("数据库关闭失败")
	//	}
	//}(database)
}
