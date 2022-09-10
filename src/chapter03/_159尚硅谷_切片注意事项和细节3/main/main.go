package main

import "fmt"

/**
日期:2022/5/14  时间:17:36
@author:冰菓春物咲太实教
*/

func main() {

	arr01 := "hmzl@qq.com雪乃"
	fmt.Println("==>", arr01[4:])
	//slice01 := []byte(arr01)
	slice01 := []rune(arr01)
	fmt.Println("==>", string(slice01))
	slice01[0] = '比'
	fmt.Println("==>", string(slice01))

}
