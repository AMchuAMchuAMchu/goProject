package main

import "fmt"

/**
日期:2022/5/13  时间:19:38
@author:冰菓春物咲太实教
*/

func main() {

	slice01 := []string{"雪之下雪乃", "堀北铃音", "千反田爱瑠", "樱岛麻衣"}
	fmt.Println("1.0 直接for循环遍历切片:")
	//fmt.Println("==>:",slice01[4])
	for i := 0; i < len(slice01); i++ {
		fmt.Println("==>", slice01[i])
	}
	fmt.Println("2.0 直接range循环遍历切片:")
	for _, s := range slice01 {
		fmt.Println("==>:", s)
	}

}
