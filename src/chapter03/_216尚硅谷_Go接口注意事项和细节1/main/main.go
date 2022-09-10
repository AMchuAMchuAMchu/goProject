package main

import "fmt"

/**
日期:2022/6/4  时间:7:12
@author:冰菓春物咲太实教
*/

type Interface interface {
	Call()
}

type Phone struct {
}

func (p Phone) Call() {
	fmt.Println("我说...")
}

//type Connect interface {
//	Connect() bool
//}
//
//func (p Phone)Connect(pwd int) bool {
//	if pwd == 3344 {
//		return true
//	}else {
//		return false
//	}
//}

func main() {
	//var p1 Phone
	var p2 Phone
	var t3 Interface = p2
	t3.Call()
	var t4 interface{} = p2
	fmt.Println(t4)
	//p1.Call()
	//var v2 Connect = p2
	//fmt.Println(p1.Connect(3344))
	//v2.Connect()
}
