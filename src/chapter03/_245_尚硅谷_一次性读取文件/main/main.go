package main

import (
	"fmt"
	"io/ioutil"
)

/**
日期:2022/6/27  时间:17:17
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {
	data, err := ioutil.ReadFile("D:\\seldom\\rd\\Go_ProjectAll\\goProject\\src\\chapter03\\assets\\image\\_01jm.txt")
	if err != nil {
		fmt.Println("!出错啦!!(•́へ•́╬)")
	}
	line := 0
	fmt.Println("第", line, "行::", string(data))
	//fmt.Println("第",line,"行::",*((*string)(unsafe.Pointer(&data))))
}
