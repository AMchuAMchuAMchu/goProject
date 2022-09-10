package main

import "fmt"

/**
日期:2022/7/7  时间:19:04
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func add() {
	for i := 0; i < 10; i++ {
		fmt.Println("add()==>", i)
	}
}

func main() {
	go add()
	for i := 0; i < 10; i++ {
		fmt.Println("main()==>", i)
	}

}
