package main

import "fmt"

/**
日期:2022/5/19  时间:19:29
@author:冰菓春物咲太实教
*/
type Cat struct {
	age   int
	color string
}

func main() {
	Cats := make(map[string]Cat)
	cat01 := Cat{
		age:   3,
		color: "白色",
	}
	cat02 := Cat{
		age:   100,
		color: "花色",
	}
	Cats["小白"] = cat01
	Cats["小白"] = cat02
	name := ""
label:
	fmt.Println("请输入猫的名字^_^:")
	fmt.Scan(&name)
	for s, _ := range Cats {
		if s != name {
			fmt.Println("!!果咩!~~o(╥﹏╥)o张老太没有这只猫...")
			goto label
		}
	}
	fmt.Println(Cats[name])

}
