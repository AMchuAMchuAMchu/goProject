package main

import "fmt"

/**
日期:2022/5/19  时间:18:36
@author:冰菓春物咲太实教
*/

type Animation struct {
	name    string
	time    int
	address string
}

func main() {

	Animations := make(map[string]Animation)
	a01 := Animation{
		name:    "鬼灭之刃",
		time:    2019,
		address: "Japanese",
	}

	a02 := Animation{
		name:    "冰海战记",
		time:    2019,
		address: "Japanese",
	}
	Animations["第一部"] = a01
	Animations["第二部"] = a02
	for s, animation := range Animations {
		fmt.Println(s, animation)
		fmt.Println(animation.name)
		fmt.Println(animation.time)
		fmt.Println(animation.address)
	}

}
