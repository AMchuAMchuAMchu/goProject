package main

import "fmt"

/**
日期:2022/4/5  时间:15:14
@author:冰菓春物咲太实教
*/

func main() {

label1:
	for i := 0; i < 4; i++ {

		for j := 0; j < 10; j++ {

			if j == 2 {
				break label1
			}
			fmt.Println("j :", j)

		}
	}

}
