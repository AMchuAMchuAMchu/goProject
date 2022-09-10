package main

import (
	"fmt"
)

/**
日期:2022/6/8  时间:6:22
@author:冰菓春物咲太实教
*/

func main() {
	score := 0
	fmt.Println("请输入您的成绩^_^:")
	fmt.Scan(&score)
	if score >= 60 {
		fmt.Println("及格...哦咩爹多~^_^")
	} else {
		fmt.Println("果咩o(╥﹏╥)o不及格未过....")
	}

}
