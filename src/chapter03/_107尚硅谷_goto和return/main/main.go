package main

import "fmt"

/**
日期:2022/4/6  时间:17:39
@author:冰菓春物咲太实教
*/
func hello() string {
	return "哦哈呦~~(≧∇≦)ﾉ"
}
func main() {
	num1 := -90
	fmt.Println("ok1")
	if num1 > 0 {
		goto label01
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	return
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
label01:
	fmt.Println("ok8")
	fmt.Println("====================================================")
	fmt.Println(hello())

}
