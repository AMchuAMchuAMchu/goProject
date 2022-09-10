package main

import (
	"fmt"
	"time"
)

/**
日期:2022/5/24  时间:19:32
@author:冰菓春物咲太实教
*/
//
//type MethodUtils struct {
//
//}
//
//func (m MethodUtils)Rectangle()  {
//	for i := 0; i < 10; i++ {
//		for j := 0; j < 8; j++ {
//			fmt.Print("✨")
//		}
//		fmt.Println()
//	}
//}

func main() {

	//MethodUtils{}.Rectangle()

	start := time.Now()
	for i := 1; i <= 1000000; i++ {
		fmt.Println(i)
	}
	end := time.Now()
	fmt.Println("耗时:", end-start)

}
