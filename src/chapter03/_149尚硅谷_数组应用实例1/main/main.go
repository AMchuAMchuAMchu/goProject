package main

import (
	"fmt"
	"time"
)

/**
日期:2022/5/12  时间:16:28
@author:冰菓春物咲太实教
*/

//func main() {
//	arr := [26]byte{}
//	//var var01 byte ='A'
//	//var end byte ='Z'
//	//arr[0] = 1
//	//fmt.Println("==>",var01)
//	//fmt.Println("==>",end)
//	var count byte = 65
//	for i := 0; i < 26; i++ {
//		arr[i] = count
//		count++
//	}
//
//	for i, b := range arr {
//		fmt.Println("下标:==>",i,"值:==>",b)
//		fmt.Println("对应字母为:",string(b))
//	}
//
//	fmt.Println("数组的最大值:",arr[len(arr)-1],"下标:",len(arr)-1)
//
//	sum := 0
//	avg :=0
//	for _, b := range arr {
//		sum+=int(b)
//		avg = sum/len(arr)
//
//	}
//		fmt.Println("和:",sum,"平均值:",avg)
//
//}

func main() {

	start := time.Now()
	for i := 1; i <= 1000000; i++ {
		fmt.Println(i)
	}
	time := time.Since(start)
	fmt.Println("耗时:", time, "毫秒.")

}
