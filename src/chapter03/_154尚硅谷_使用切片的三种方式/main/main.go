package main

import "fmt"

/**
日期:2022/5/12  时间:18:05
@author:冰菓春物咲太实教
*/

func main() {
	var slice01 []float64
	//fmt.Printf("slice01:%T",slice01)
	slice02 := make([]string, 4, 10)
	slice03 := make([]float64, 4, 10)
	fmt.Println("::==>", slice01)
	fmt.Println("::==>", slice02)
	fmt.Println("::==>", slice03)
	slice02[0] = "罪恶王冠"
	slice02[1] = "荒木折郎"
	slice02[2] = "樱满集"
	slice02[3] = "蝶祈"
	fmt.Println("打印:", slice02)
	fmt.Println("长度:", len(slice02))
	fmt.Println("容量:", cap(slice02))
	var slice04 = []int{11, 111, 222, 333}
	fmt.Println("===>", slice04)

}
