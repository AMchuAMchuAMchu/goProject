package main

import (
	"chapter03/_205尚硅谷_封装快速入门案例/person"
	"fmt"
)

/**
日期:2022/6/1  时间:21:50
@author:冰菓春物咲太实教
*/

func main() {
	p1 := person.NewPerson("雪乃")
	fmt.Println(p1.GetAge())
	p1.SetAge(34)
	fmt.Println(p1.GetAge())
}
