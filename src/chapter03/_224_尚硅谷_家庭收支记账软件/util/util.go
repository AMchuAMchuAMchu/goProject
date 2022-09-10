package util

import (
	"fmt"
	"strconv"
)

/**
日期:2022/6/24  时间:7:34
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func (this *userList) Menu() {

	times := 0
outer:
	for {
		fmt.Println(">>>>>>>>>>>>>>>>>>家庭收支记账软件<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退   出")
		fmt.Println()
		fmt.Print("请客官选择(1-4)^_^:")
		var option int64 = 4
		input := ""
		fmt.Scanln(&input)
		option, _ = strconv.ParseInt(input, 10, 64)
		//list := userList{}
		//fmt.Println("==>",list.GetUserList)
		switch option {
		case 1:
			//f_IAE()
			this.f_IncomeAndExpenses()
			fmt.Println()
		case 2:
			//f_I()
			this.f_income()
			fmt.Println()
		case 3:
			//f_e()
			this.f_expenses()
			fmt.Println()
		case 4:
			key := ""
			fmt.Print("确定退出吗?o(╥﹏╥)o(确定退出请输入:Y/y)")
			fmt.Scanln(&key)
			if key == "Y" || key == "y" {
				fmt.Println("撒由那拉...o(╥﹏╥)o😭😭😭")
				break outer
			}
			fmt.Println()
		default:
			times++
			fmt.Println("输入异常累计", times, "次!!还剩下", 5-times, "次机会!!(•́へ•́╬)请重新输入!!!")
			fmt.Println()
			if times == 5 {
				fmt.Println("!!警告!!输入错误次数过多,为了您的账号安全将锁定账号三十年!!(•́へ•́╬)不服直接到线下营业厅办理...")
				return
			}

		}

	}
}

type user struct {
	Name              string
	Balance           float64
	IncomeAndExpenses float64
	Description       string
}

type userList struct {
	GetUserList []user
}

func (this *userList) f_IncomeAndExpenses() {
	fmt.Println("1 收支明细")
	fmt.Print("用户名\t余额\t收支\t说明")
	userListInfo := this.GetUserList
	//fmt.Println(userListInfo)
	for i := 0; i < len(userListInfo); i++ {
		fmt.Printf("\n%v\t%v\t%v\t%v",
			userListInfo[i].Name, userListInfo[i].Balance, userListInfo[i].IncomeAndExpenses,
			userListInfo[i].Description)
	}
}

func (this *userList) f_income() {
	name := ""
	balance := 20000.0
	incomeAndExpenses := 0.0
	description := ""
	fmt.Println("2 登记收入")
	fmt.Print("^_^请输入用户名:")
	fmt.Scanln(&name)
	fmt.Print("^_^请输入收支:")
	fmt.Scanln(&incomeAndExpenses)
	fmt.Print("^_^请输入说明:")
	fmt.Scanln(&description)
	balance += incomeAndExpenses
	this.GetUserList = append(this.GetUserList, user{
		Name:              name,
		Balance:           balance,
		IncomeAndExpenses: incomeAndExpenses,
		Description:       description,
	})
	fmt.Println("收入记账成功(≧∇≦)ﾉ...")
	fmt.Println(this.GetUserList[0])
}

func (this *userList) f_expenses() {
	name := ""
	balance := 20000.0
	incomeAndExpenses := 0.0
	description := ""
	fmt.Println("3 登记支出")
	fmt.Print("^_^请输入用户名:")
	fmt.Scanln(&name)
	fmt.Print("^_^请输入收支:")
	fmt.Scanln(&incomeAndExpenses)
	fmt.Print("^_^请输入说明:")
	fmt.Scanln(&description)
	balance -= incomeAndExpenses
	this.GetUserList = append(this.GetUserList, user{
		Name:              name,
		Balance:           balance,
		IncomeAndExpenses: incomeAndExpenses,
		Description:       description,
	})
	fmt.Println("支出记账成功(≧∇≦)ﾉ...")
}

func NewuserList() *userList {
	return &userList{GetUserList: nil}
}
