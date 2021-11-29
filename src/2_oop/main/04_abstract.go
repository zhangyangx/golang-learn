package main

import "fmt"

// 抽象
// 多个不同账户有相同的属性、功能，存款、取款等，把他们封装到同一个结构体中
func main() {
	account := Account{AccountNo: "123456", Pwd: "123", Balance: 0.0}
	fmt.Println(account)
	account.saveMoney(20)
	fmt.Println(account.Balance)
	msg := account.checkCAccount(account.AccountNo, account.Pwd)
	fmt.Println(msg)
}

// Account 定义一个账户的结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款
func (account *Account) saveMoney(money float64) {
	account.Balance = account.Balance + money
}

// 取款
func (account *Account) withdrawMoney(money float64) {
	account.Balance = account.Balance - money
}

// 校验账号密码
func (account Account) checkCAccount(accountNo string, pwd string) (msg string) {
	// 模拟
	if len(accountNo) < 1 {
		return "输入错误"
	}
	if len(pwd) < 1 {
		return "输入错误"
	}
	return "success"
}
