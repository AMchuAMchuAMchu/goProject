package main

import (
	"chapter03/_201尚硅谷_工厂模式详解/model"
	"fmt"
)

/**
日期:2022/5/27  时间:17:59
@author:冰菓春物咲太实教
*/

func main() {

	//fmt.Println(model.Animal{
	//	Name: "dog...",
	//	Age:  2,
	//})
	m01 := model.GetAnimalInstance("cat...", 11)
	//fmt.Printf("%T",m01)
	fmt.Println(m01)

}
