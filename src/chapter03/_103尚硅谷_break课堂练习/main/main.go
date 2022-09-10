package main

import "fmt"

/**
日期:2022/4/6  时间:15:27
@author:冰菓春物咲太实教
*/

func main() {
	//sum := 0
	//for i := 1;i <= 100;i++ {
	//	sum += i
	//	if sum > 20 {
	//		fmt.Println("当前数:",i)
	//		break
	//	}
	//}

	user := "张无忌"
	password := "888"
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名:")
		fmt.Scanln(&user)
		fmt.Println("请输入密码:")
		fmt.Scanln(&password)
		if user == "张无忌" && password == "888" {
			fmt.Println("哦咩爹度!!~(≧∇≦)ﾉ登陆成功!!")
			break

		} else {
			fmt.Println("很遗憾....o(╥﹏╥)o,登陆失败!!你还剩下", 3-i, "次机会")
			if i == 3 {
				fmt.Println("残念....(⊙︿⊙)..直接结束登录,请您明天再来...^_^")
				break
			}
		}

	}

}
