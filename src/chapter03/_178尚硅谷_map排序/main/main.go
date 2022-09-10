package main

import "fmt"

/**
日期:2022/5/19  时间:18:14
@author:冰菓春物咲太实教
*/

func main() {

	//slice_index01 := []int{0,1,2,3,4}
	//map01 := map[int]string{"0":"乃","1":"下","2":"之","3":"雪"}
	map01 := make(map[int]string)
	map01[0] = "雪"
	map01[1] = "之"
	map01[2] = "下"
	map01[3] = "雪"
	map01[4] = "乃"
	fmt.Println(map01)
	//直接for range输出也是可以实现有序输出????
	//for i, s := range map01 {
	//	fmt.Println(i,s)
	//}
	//for i := 4; i >= 0; i-- {
	//	fmt.Println(map01[i])
	//}

}
