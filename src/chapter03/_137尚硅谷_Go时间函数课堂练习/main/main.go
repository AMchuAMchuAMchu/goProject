package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
日期:2022/5/6  时间:14:20
@author:冰菓春物咲太实教
*/

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {

	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println("耗时:", end-start)
}
