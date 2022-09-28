package main

import "fmt"

/**
日期:2022/4/5  时间:14:17
@author:冰菓春物咲太实教
*/
//判断100 - 1000 范围的水仙花数

func main() {

	for i := 100; i <= 1000; i++ {

		n1 := i / 100
		n2 := i % 100 / 10
		n3 := i % 10
		if n1*n1*n1+n2*n2*n2+n3*n3*n3 == i {
			fmt.Println(i)
		}

		n4 := i / 100
		n5 := i % 100 / 10
		n6 := i % 10
		if n4*n4*n4+n5*n5*n5+n6*n6*n6 == i {
			fmt.Println(i)
		}

	}

}
