package main

import (
	"fmt"
	"strconv"
)

/**
日期:2022/5/5  时间:11:12
@author:冰菓春物咲太实教
*/

func main() {

	str := "hellogolang"
	str01 := "hellogolang雪乃"
	fmt.Println("str 的长度:", len(str))
	fmt.Println("str 的长度:", len(str01))
	res01 := []rune(str01)
	for i := 0; i < len(res01); i++ {
		fmt.Printf("输出:%c||", res01[i])
	}
	fmt.Println()
	fmt.Println("=================================")
	n, err := strconv.Atoi("996")
	if err != nil {
		fmt.Println("转换错误!!(•́へ•́╬)")
	} else {
		fmt.Println("转换正确!!!(≧∇≦)ﾉ:", n)
	}
	istr := strconv.Itoa(9651314)
	fmt.Printf("str:%v,type:%T", istr, istr)

}
