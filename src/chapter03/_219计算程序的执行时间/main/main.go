package main

import (
	"fmt"
	"time"
)

/**
日期:2022/6/8  时间:6:22
@author:冰菓春物咲太实教
*/

func main() {
	start := time.Now()
	count := 0
	for i := 1; i <= 1000000000; i++ {
		//fmt.Println(i)
		count += i
	}
	cTime := time.Since(start)
	fmt.Println("耗时:", cTime)
	/*
		第一次:耗时: 253.076ms
		第二次:耗时: 250.0202ms
		第三次:耗时: 243.1419ms
	*/
}
