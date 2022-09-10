package main

import (
	"encoding/json"
	"fmt"
)

/**
日期:2022/7/2  时间:18:30
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

type User struct {
	Username string
	Password string
	Sex      string
	Hobby    string
	Travel   string
}

func main() {
	instance_user := &User{
		Username: "刀剑神域",
		Password: "亚丝娜",
		Sex:      "女",
		Hobby:    "玩刀剑游戏",
		Travel:   "SOA",
	}
	res, err := json.Marshal(instance_user)
	if err != nil {
		fmt.Println("!!果咩~~出现异常(•́へ•́╬)!")
		return
	}
	fmt.Println(string(res))
}
