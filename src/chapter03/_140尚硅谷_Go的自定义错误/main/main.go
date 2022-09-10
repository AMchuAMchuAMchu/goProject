package main

import (
	"errors"
	"fmt"
	"time"
)

/**
日期:2022/5/8  时间:20:18
@author:冰菓春物咲太实教
*/

func readFile(path string) (err error) {
	if path == "config.txt" {
		return nil
	} else {
		return errors.New("文件读取错误!!!..(•́へ•́╬)")
	}
}

func test01() {
	res := readFile("config.tx")
	if res != nil {
		panic(res)
	}

	fmt.Println("test01继续执行....")

}

func main() {
	test01()
	for {

		fmt.Println("mian继续执行....")
		time.Sleep(time.Second)
	}

}
