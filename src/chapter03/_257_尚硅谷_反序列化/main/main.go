package main

import (
	"encoding/json"
	"fmt"
)

/**
日期:2022/7/3  时间:21:13
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
	//u1 := User{
	//	Username: "梅诺",
	//	Password: "5203344",
	//	Sex:      "女",
	//	Hobby:    "我是一个清廉正直且强大的神官",
	//	Travel:   "和时任灯里....",
	//}
	//
	//res,_ := json.Marshal(u1)
	//fmt.Println(string(res))

	u2 := "{\"Username\":\"梅诺\",\"Password\":\"5203344\"," +
		"\"Sex\":\"女\",\"Hobby\":\"我是一个清廉正直且强大的神官\"," +
		"\"Travel\":\"和时任灯里....\"}"

	var u3 User

	err := json.Unmarshal([]byte(u2), &u3)
	if err != nil {
		fmt.Println("!!!")
		return
	}
	fmt.Println(u3)

}
