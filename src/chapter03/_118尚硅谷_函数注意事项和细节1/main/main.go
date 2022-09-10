package main

import "fmt"

/**
日期:2022/4/15  时间:15:09
@author:冰菓春物咲太实教
*/

func add(num1 *int) int {
	//*num1 = *num1 + 10
	//return *num1
	return *num1 + 10
}

func main() {
	num1 := 20
	fmt.Println("add:", add(&num1), "原来:", num1)
}
