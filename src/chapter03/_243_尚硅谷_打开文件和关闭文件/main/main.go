package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
日期:2022/6/25  时间:21:26
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {

	file, err := os.Open("D:\\seldom\\rd\\goProject\\src\\chapter03\\assets\\image\\_01jm.txt")
	//if err != nil {
	//	fmt.Println("完了!!文件读取出错啦(•́へ•́╬)!!")
	//}
	//
	//if file != nil {
	//	fmt.Println("1.0 ==>",file)
	//	fmt.Printf("2.0 ==>%v",file)
	//	fmt.Println("3.0 ==>",file)
	//}
	//fmt.Println()
	//err01 := file.Close()
	//fmt.Println("3.0 ==>",err01)

	if err != nil {
		fmt.Println("!!警告!!(•́へ•́╬)发生错误地说...")
	}

	defer file.Close()

	fileReader := bufio.NewReader(file)
	line := 0
	for {
		line++
		str, err01 := fileReader.ReadString('\n')
		if err01 != nil {
			fmt.Println("!!警告!!(•́へ•́╬)发生错误地说...==>", err01)
			break
		}
		fmt.Print("第", line, "行::", str)
	}

}
