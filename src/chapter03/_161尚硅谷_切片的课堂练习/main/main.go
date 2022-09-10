package main

import "fmt"

/**
日期:2022/5/14  时间:17:55
@author:冰菓春物咲太实教
*/

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 0
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {

	//fmt.Println("==>",fbn(20))

	var slice01 []int = []int{1, 2, 3, 4, 5}
	var slice02 = make([]int, 10)
	slice01[0] = 99
	copy(slice02, slice01)
	fmt.Println("slice1=", slice01)
	fmt.Println("slice2=", slice02)

}
