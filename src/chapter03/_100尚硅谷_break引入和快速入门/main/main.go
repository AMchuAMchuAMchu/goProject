package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
日期:2022/4/5  时间:14:28
@author:冰菓春物咲太实教
*/

func main() {

	count := 0
	n1 := 0
	//rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())
	for {
		n1 = rand.Intn(100)
		count++
		//fmt.Println("打印:",n1)
		fmt.Println(n1)
		if n1 == 99 {
			fmt.Println("99找到了!!(p≧w≦q),一共随机生成了", count, "次!")
			break
		}
	}
	//os.Exit(0)
	fmt.Println("撒由那拉~~o(╥﹏╥)o(⊙︿⊙)")

}
