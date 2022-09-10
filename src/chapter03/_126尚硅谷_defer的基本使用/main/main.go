package main

import "fmt"

/**
日期:2022/4/28  时间:17:15
@author:冰菓春物咲太实教
*/

func _defer(num1 int, num2 int) int {

	defer fmt.Println("输出num1:", num1)
	defer fmt.Println("输出num2:", num2)
	res := num1 + num2
	fmt.Println("输出结果:", res)
	return res
}

func main() {

	res01 := _defer(234, 567)
	fmt.Println("在main方法里面输出:", res01)

}
