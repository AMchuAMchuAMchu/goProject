package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
日期:2022/6/27  时间:17:08
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {

	fileObj, err := os.Open("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_01jm.txt")
	if err != nil {
		fmt.Println("!!警告(•́へ•́╬)程序出错要没~~~")
	}

	defer fileObj.Close()

	fileObj01 := bufio.NewReader(fileObj)
	line := 0
	for {
		line++
		info, err01 := fileObj01.ReadString('\n')

		if err01 != nil {
			fmt.Println("==>出错啦!!(•́へ•́╬)")
			break
		}

		fmt.Print("第", line, "行::", info)
	}
}
