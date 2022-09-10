package main

import "fmt"

/**
日期:2022/5/31  时间:12:14
@author:冰菓春物咲太实教
*/

type Account struct {
	AccountId  int
	AccountPwd int
	Balance    float64
}

//存钱
func (account *Account) Deposit(AccountId int, AccountPwd int, money float64) {

	if AccountId != account.AccountId {
		fmt.Println("账号错误!!(•́へ•́╬)")
		return
	} else if AccountPwd != account.AccountPwd {
		fmt.Println("密码错误!!(•́へ•́╬)")
		return
	} else if money < 0 {
		fmt.Println("金额输入异常!!(•́へ•́╬)")
		return
	}
	account.Balance += money
	fmt.Println("!(≧∇≦)ﾉ存款成功!!")

}

//取钱
func (account *Account) WithDraw(AccountId int, AccountPwd int, money float64) float64 {

	if AccountId != account.AccountId {
		fmt.Println("账号错误!!(•́へ•́╬)")
		return -1
	} else if AccountPwd != account.AccountPwd {
		fmt.Println("密码错误!!(•́へ•́╬)")
		return -1
	} else if money < 0 {
		fmt.Println("金额输入异常!!(•́へ•́╬)")
		return -1
	}
	if money > account.Balance {
		fmt.Println("Error!!超出余额!!(•́へ•́╬)")
		return -1
	}
	fmt.Println("!(≧∇≦)ﾉ存款成功!!")
	account.Balance -= money
	return money
}

func (account *Account) QueryBalance() float64 {
	return account.Balance
}

func main() {

	xnAccount := Account{
		AccountId:  5201314,
		AccountPwd: 520,
		Balance:    520,
	}
	//fmt.Println("存款前:",xnAccount.QueryBalance())
	//xnAccount.Deposit(5201314,520,1480)
	//fmt.Println("存款后:",xnAccount.QueryBalance())

	//fmt.Println("取款前:",xnAccount.Balance)
	fmt.Println("取款前:", xnAccount.QueryBalance())
	re := xnAccount.WithDraw(5201314, 520, 480)
	//fmt.Println("取款后:",xnAccount.Balance)
	fmt.Println("取款后:", xnAccount.QueryBalance())
	fmt.Println("取出:", re)
}
