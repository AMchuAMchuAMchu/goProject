package main

import "fmt"

/**
日期:2022/5/15  时间:19:49
@author:冰菓春物咲太实教
*/

func main() {

	//td := [4][6]int{}
	var td [4][6]int = [4][6]int{}
	td[1][2] = 1
	td[2][1] = 2
	td[2][3] = 3
	fmt.Println(td)
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(td[i][j], " ")
		}
		fmt.Println()
	}

}
