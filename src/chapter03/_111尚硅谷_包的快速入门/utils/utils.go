package utils

import "fmt"

/**
日期:2022/4/8  时间:10:00
@author:冰菓春物咲太实教
*/

/**
计算加减乘除的一个通用的方法!!外部用户输入参数的形式(•́へ•́╬)
*/
func Calc(num1 float64, num2 float64, str string) float64 {
	var res float64
	switch str {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		fmt.Println("输入异常!!请重新输入!!(•́へ•́╬)")
		res = 0.0
	}
	return res
}

/**
计算加减乘除的一个通用的方法!!外部用户不输入参数的形式,直接调用即可(•́へ•́╬)
*/
func CalcClient() {
	var _num1 float64
	var _num2 float64
	var _cal string
	fmt.Println("请客官输入第一个数^_^:")
	fmt.Scanln(&_num1)
	fmt.Println("请客官输入第二个数^_^:")
	fmt.Scanln(&_num2)
	fmt.Println("请客官输入运算符^_^:")
	fmt.Scanln(&_cal)
	result := Calc(_num1, _num2, _cal)
	fmt.Println("结果输出:", result)
	var res float64
	switch _cal {
	case "+":
		res = _num1 + _num2
	case "-":
		res = _num1 - _num2
	case "*":
		res = _num1 * _num2
	case "/":
		res = _num1 / _num2
	default:
		fmt.Println("输入异常!!请重新输入!!(•́へ•́╬)")
		res = 0.0
	}
	fmt.Println("久等啦!!(≧∇≦)ﾉ输出的结果是:", res)
}
