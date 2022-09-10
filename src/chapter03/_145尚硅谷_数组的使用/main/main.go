package main

import "fmt"

/**
日期:2022/5/9  时间:18:20
@author:冰菓春物咲太实教
*/

func arrFloat() {
	var arrFloat [5]float64
	count := 0
	for i := 0; i < len(arrFloat); i++ {
		count++
		fmt.Println("请客官输入第", count, "个数^_^:")
		fmt.Scan(&arrFloat[i])
	}
	fmt.Println("==>", arrFloat)
	fmt.Printf("类型:%T", arrFloat)

}

func main() {
	//arrFloat()

	var arr01 [3]int
	var arr02 [3]int = [3]int{1, 1, 1}
	var arr03 = [3]int{1, 1, 1}
	var arr04 = [...]int{1, 1, 1, 2}
	fmt.Println("==>", arr01)
	fmt.Println("==>", arr02)
	fmt.Println("==>", arr03)
	fmt.Println("==>", arr04)

}
