package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
日期:2022/5/12  时间:17:18
@author:冰菓春物咲太实教
*/

func main() {

	rand.Seed(time.Now().UnixNano())
	arr := [5]int{}
	arr01 := [5]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("反转之前:")
	for _, i := range arr {
		fmt.Println(i)
	}

	fmt.Println("反转之后:")
	for i := 4; i >= 0; i-- {
		arr01[len(arr)-1-i] = arr[i]

	}
	arr = arr01

	for _, i := range arr {
		fmt.Println(i)
	}

}
