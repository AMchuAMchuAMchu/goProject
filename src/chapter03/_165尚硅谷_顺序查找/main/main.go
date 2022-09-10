package main

import "fmt"

/**
日期:2022/5/15  时间:17:39
@author:冰菓春物咲太实教
*/

func main() {
	slice01 := []string{"雪乃", "由比滨结衣", "樱岛麻衣", "磷叶石"}
	name := ""
label:
	for i := 0; i < 5; i++ {

		fmt.Println("请客官输入:")
		fmt.Scan(&name)
		flag := true
		for i := 0; i < len(slice01); i++ {
			if slice01[i] == name {
				flag = false
				fmt.Println("哦咩爹多~~恭喜你找到了!!^_^")
				break label
			}

		}
		if flag {
			fmt.Println("果咩捏~~抱歉~~o(╥﹏╥)o没能找到...")
		}

	}
}
