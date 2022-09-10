package main

import "fmt"

/**
日期:2022/5/22  时间:15:29
@author:冰菓春物咲太实教
*/

type Person struct {
	Name string
}

func (p Person) call() {
	fmt.Println("call==>比企谷快速过来!!!(•́へ•́╬)...")
}

func main() {
	p1 := Person{Name: "jjj"}
	p1.call()
}
