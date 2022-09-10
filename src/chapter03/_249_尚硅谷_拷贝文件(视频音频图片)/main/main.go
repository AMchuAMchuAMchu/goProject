package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
日期:2022/6/27  时间:18:34
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {
	data, err := os.Open("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_06jm.jpg")
	if err != nil {
		fmt.Println("出现错误!!(•́へ•́╬)")
	}
	defer data.Close()
	data01 := bufio.NewReader(data)

	data03, err01 := os.OpenFile("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_07jmcopy.jpg",
		os.O_WRONLY|os.O_CREATE, 666)
	defer data03.Close()
	if err01 != nil {
		fmt.Println("!!警告!!报错!!(•́へ•́╬)")
	}

	data04 := bufio.NewWriter(data03)

	line, err03 := io.Copy(data04, data01)
	if err03 != nil {
		fmt.Println("报错!!==>", err03)
	}
	fmt.Println(line)

}
