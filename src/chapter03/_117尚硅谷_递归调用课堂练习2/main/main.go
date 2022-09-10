package main

import "fmt"

/**
日期:2022/4/11  时间:18:35
@author:冰菓春物咲太实教
*/

//func fbn(num1 int64) int64 {
//
//	if num1 == 1 || num1 == 2 {
//		return 1
//	}else {
//		return fbn(num1 - 1) + fbn(num1 - 2)
//	}
//
//}
//
//func main() {
//
//	i := fbn(4)
//	i2 := fbn(5)
//	fmt.Println(i)
//	fmt.Println(i2)
//
//}

func peach(day int) int {
	if !(day >= 1 || day <= 10) {
		return -1
	}
	if day == 10 {
		return 1
	} else {
		return (peach(day+1) + 1) * 2
	}

}

func main() {
	fmt.Println("第一天的桃子数目:", peach(1))
}
