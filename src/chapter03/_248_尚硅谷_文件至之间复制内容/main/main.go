package main

import (
	"fmt"
	"io/ioutil"
)

/**
日期:2022/6/27  时间:18:16
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {
	fileInfo, err := ioutil.ReadFile("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_03ysn.txt")
	if err != nil {
		fmt.Println("::", err)
	}
	err01 := ioutil.WriteFile("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_04.ysn.txt", fileInfo, 888)
	if err01 != nil {
		fmt.Println("出错啦!!(•́へ•́╬)")
	}
}
