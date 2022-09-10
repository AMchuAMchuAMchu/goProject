package main

import "fmt"

/**
日期:2022/5/14  时间:16:03
@author:冰菓春物咲太实教
*/

func main() {

	var slice01 []string = []string{"锁那", "yoasobi", "Aimer", "milet"}
	var slice02 []string = []string{"樱满集", "蝶祈", "罪恶王冠", "荒木哲郎"}
	slice03 := append(slice02, "筱宫绫濑", "神之塔")
	slice04 := append(slice03, slice01...)
	fmt.Println("slice01==>:", slice01)
	fmt.Println("slice02==>:", slice02)
	fmt.Println("slice03==>:", slice03)
	fmt.Println("slice03==>:", slice04)

}
