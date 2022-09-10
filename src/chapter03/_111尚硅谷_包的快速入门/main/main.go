package main

import (
	u "chapter03/_111尚硅谷_包的快速入门/utils"
	"fmt"
)

/**
日期:2022/4/8  时间:10:00
@author:冰菓春物咲太实教
*/

func main() {

	fmt.Println("Calc方法调用测试:")
	//utils.Calc(34,78,"+")
	u.Calc(34, 78, "+")
	fmt.Println("CalcClient方法调用测试:")
	//utils.CalcClient()
	u.CalcClient()

}
