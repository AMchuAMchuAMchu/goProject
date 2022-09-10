package main

import "fmt"

/**
日期:2022/5/10  时间:20:12
@author:冰菓春物咲太实教
*/

func main() {

	arr01 := [...]string{"雪之下", "由比滨结衣", "鸭志田一", "青猪"}
	for i, s := range arr01 {
		fmt.Println("索引:", i, "值:", s)
	}

}
