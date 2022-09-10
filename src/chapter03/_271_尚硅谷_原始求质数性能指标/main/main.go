package main

import (
	"fmt"
	"time"
)

/**
*@BelongsProject: goProject
*@BelongsPackage:
*@Author: 02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
*@CreateTime: 2022-07-24 13:00:57
*@Description: TODO
*@Version: 1.0
 */

func main() {
	arr := make([]int, 0)
	var t int
	start := time.Now().UnixMilli()
	for i := 2; i <= 80000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag {
			t++
			//fmt.Println(i)
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
	end := time.Now().UnixMilli() //耗时:: 2011 毫秒,个数: 7838
	fmt.Println("耗时::", end-start, "毫秒,个数:", t)
	//耗时:: 2011 毫秒,个数: 7838    耗时:: 1991 毫秒,个数: 7838   耗时:: 2010 毫秒,个数: 7838
}
