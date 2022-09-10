package main

import "fmt"

/**
日期:2022/5/15  时间:15:20
@author:冰菓春物咲太实教
*/

func main() {

	slice01 := []int{12, 56, 7, 99, 34, 47, 899, 455, 89899, 9}
	fmt.Println("冒泡排序之前(升序)切片顺序:", slice01)
	for i := len(slice01) - 1; i > 0; i-- {
		for j := 0; j < len(slice01)-1; j++ {
			if slice01[j+1] < slice01[j] {
				tem := slice01[j+1]
				slice01[j+1] = slice01[j]
				slice01[j] = tem
			}
		}
	}
	fmt.Println("冒泡排序之后(升序)切片顺序:", slice01)

	fmt.Println("冒泡排序之前(降序)切片顺序:", slice01)
	for i := len(slice01) - 1; i > 0; i-- {
		for j := 0; j < len(slice01)-1; j++ {
			if slice01[j+1] > slice01[j] {
				tem := slice01[j+1]
				slice01[j+1] = slice01[j]
				slice01[j] = tem
			}
		}
	}
	fmt.Println("冒泡排序之后(降序)切片顺序:", slice01)

}
