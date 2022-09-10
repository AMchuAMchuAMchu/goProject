package main

import "fmt"

/**
日期:2022/5/27  时间:17:44
@author:冰菓春物咲太实教
*/

type Animal struct {
	Name string
	Age  int
}

func main() {
	a01 := Animal{
		Name: "dog",
		Age:  1,
	}
	fmt.Println(a01)

}
