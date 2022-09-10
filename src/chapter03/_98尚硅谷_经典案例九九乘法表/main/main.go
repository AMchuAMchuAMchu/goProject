package main

import "fmt"

/**
日期:2022/4/5  时间:14:09
@author:冰菓春物咲太实教
*/

func main() {

	num := 0
	fmt.Println("请输入:")
	fmt.Scanln(&num)

	for r := 1; r <= num; r++ {

		for c := 1; c <= r; c++ {
			fmt.Print(r, " * ", c, " = ", r*c, "\t")
		}
		fmt.Println()
	}

}
