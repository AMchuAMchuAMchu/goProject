package main

import "fmt"

/**
日期:2022/5/19  时间:18:51
@author:冰菓春物咲太实教
*/

func addUser() {
	users := make(map[string]map[string]string)
	users["1"] = map[string]string{"nickname": "蝶祈", "pwd": "912"}
	users["2"] = map[string]string{"nickname": "樱满集", "pwd": "821"}
	users["3"] = map[string]string{"nickname": "恙神涯", "pwd": "820"}
	name := ""
	nickname := ""
	pwd := ""
label:
	fmt.Println("请输入您需要寻找的角色^_^:")
	fmt.Scan(&name)
	for k, _ := range users {
		if name == k {
			fmt.Println("!!!(•́へ•́╬)用户名已重复!!!请重新输入谢谢(≧∇≦)ﾉ!!!:")
			//todo ==> 直接使用方法递归/goto语句
			goto label
		}
	}
	fmt.Println("请输入用户昵称^_^:")
	fmt.Scan(&nickname)
	fmt.Println("请输入用户密码^_^:")
	fmt.Scan(&pwd)
	users[name] = map[string]string{"nickname": nickname, "pwd": pwd}
	fmt.Println("添加完毕!!(≧∇≦)ﾉ")
	fmt.Println("添加之后的结果:", users)
}

func main() {
	addUser()
}
