package main

import "fmt"

/**
日期:2022/4/6  时间:16:01
@author:冰菓春物咲太实教
*/

func main() {
	//var 雪乃 int;
	//fmt.Println(雪乃)
	//for i := 0;i <= 100;i++ {
	//	if i % 2 == 0{
	//		continue
	//	}
	//	fmt.Println("输出奇数:",i)
	//
	//
	//}

	positive := 0
	negative := 0
	input := 0
	for {
		fmt.Println("请输入:")
		fmt.Scanln(&input)
		if input > 0 {
			positive++
			continue
		} else if input == 0 {
			break
		}
		negative++

	}
	fmt.Println("正数个数:", positive)
	fmt.Println("负数个数:", negative)

}
