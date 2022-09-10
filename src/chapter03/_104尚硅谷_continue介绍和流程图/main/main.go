package main

import "fmt"

/**
日期:2022/4/6  时间:15:50
@author:冰菓春物咲太实教
*/

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("输出:", i)
	}
}
