package main

import (
	"fmt"
)

/**
日期:2022/5/15  时间:18:08
@author:冰菓春物咲太实教
*/

func main() {

	slice01 := []int{2, 12, 26, 56, 78, 218, 990}
	user := 0
	counter := 0
label:
	for i := 0; i < 5; i++ {
		fmt.Println("请客官输入(≧∇≦)ﾉ:")
		fmt.Scan(&user)
		leftIndex := 0
		rightIndex := len(slice01) - 1
		for i := 0; i < len(slice01); i++ {
			counter++
			//四舍五入解决出现leftindex和rightindex永远不相等的情况.....^_^
			middle := (leftIndex + rightIndex) / 2
			if user < slice01[middle] {
				rightIndex = middle
				//到这里的话就可以直接另外分析了....
				if slice01[leftIndex] == user {
					fmt.Println("哦咩爹多~~找到了喔~^_^,是第", leftIndex+1, "个")
					break label
				}
			} else if slice01[middle] < user {
				leftIndex = middle
				//到这里的话就可以直接另外分析了....已经只能和最后一个比较了,自己分析即可....^_^
				if slice01[rightIndex] == user {
					fmt.Println("哦咩爹多~~找到了喔~^_^,是第", rightIndex+1, "个")
					break label
				}
			} else {
				fmt.Println("哦咩爹多~~找到了喔~^_^")
				break label
			}
			fmt.Println("leftIndex:", leftIndex)
			fmt.Println("rightIndex:", rightIndex)
			//fmt.Println("==>",counter)
			if leftIndex == rightIndex {
				fmt.Println("果咩纳塞~~没有找到...o(╥﹏╥)o")
				break label
			}
		}
	}

}
