package main

import (
	"fmt"
	"time"
)

/**
日期:2022/5/8  时间:20:00
@author:冰菓春物咲太实教
*/

func multi() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误信息打印:", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println("打印结果:", num1/num2)
}

func main() {

	multi()
	for {

		fmt.Println("跳过异常错误执行代码了了吗???^_^")
		time.Sleep(time.Second * 2)
	}

}
