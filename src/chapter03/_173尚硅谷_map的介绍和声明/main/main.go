package main

import "fmt"

/**
日期:2022/5/16  时间:19:05
@author:冰菓春物咲太实教
*/

func main() {

	map01 := map[string]string{}
	var map02 map[int]string
	map02 = make(map[int]string, 10)
	map01["第一部"] = "雪乃"
	map01["第二部"] = "樱岛麻衣"
	fmt.Println(map01)
	map02[1] = "666"
	map02[2] = "886"
	fmt.Println(map02)

}
