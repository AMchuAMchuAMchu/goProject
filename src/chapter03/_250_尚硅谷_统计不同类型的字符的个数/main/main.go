package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
日期:2022/6/28  时间:13:50
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

type countText struct {
	EnCount     int
	NumCount    int
	SpaceCount  int
	OthersCount int
}

func main() {

	file, err := os.Open("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\text\\_01_random.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("!!警告!!出错啦!!(•́へ•́╬)")
		return
	}

	data := bufio.NewReader(file)
	var ct countText
	for {
		textInfo, err01 := data.ReadString('\n')

		for _, v := range textInfo {
			//fmt.Println("k:",k)
			//fmt.Println("v:",v)
			switch {
			case v >= 'a' && v <= 'z':
				ct.EnCount++
			case v >= 'A' && v <= 'Z':
				ct.EnCount++
			case v >= '0' && v <= '9':
				ct.NumCount++
			case v == ' ' || v == '\t':
				ct.SpaceCount++
			default:
				ct.OthersCount++
			}
		}
		if err01 != nil {
			fmt.Println("!!警告!!(•́へ•́╬)")
			break //在循环的话就是直接break了,return的话下面的代码的话是无法到达的说....
		}

	}

	fmt.Println("==>", ct)

}
