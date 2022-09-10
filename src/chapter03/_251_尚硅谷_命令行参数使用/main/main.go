package main

import (
	"fmt"
	"os"
)

/**
日期:2022/6/28  时间:14:47
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {
	//go build -o test.exe  main.go
	//test.exe 127.0.0.1 d://localhost:3306 8080
	var arr []string
	arr = os.Args
	for k, v := range arr {
		fmt.Printf("k:%v == v:%v\n", k, v)
	}
}
