package main

import "fmt"

/**
日期:2022/4/11  时间:17:50
@author:冰菓春物咲太实教
*/

func calc(num1 float64, num2 float64) (float64, float64, float64, float64) {
	res1 := 0.0
	res2 := 0.0
	res3 := 0.0
	res4 := 0.0
	res1 = num1 + num2
	res2 = num1 - num2
	res3 = num1 * num2
	res4 = num1 / num2
	return res1, res2, res3, res4

}

func main() {
	num1 := 0.0
	num2 := 0.0
	fmt.Println("请输入第一个数:")
	fmt.Scanln(&num1)
	fmt.Println("请输入第二个数:")
	fmt.Scanln(&num2)
	sum := 0.0
	sub := 0.0
	mul := 0.0
	div := 0.0
	sum, sub, mul, div = calc(num1, num2)

	fmt.Println("+结果:", sum)
	fmt.Println("+结果:", sub)
	fmt.Println("+结果:", mul)
	fmt.Println("+结果:", div)

}
