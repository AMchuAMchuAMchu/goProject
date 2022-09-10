package main

import (
	"fmt"
	"time"
)

/**
日期:2022/7/13  时间:16:22
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func wd(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入::", i)
		//time.Sleep(time.Microsecond*20)
	}
	close(intChan)
}

func rd(intChan chan int, boolChan chan bool) {
	for v := range intChan {
		fmt.Println("读取到数据::", v)
		time.Sleep(time.Second * 2)
	}
	boolChan <- true
	close(boolChan)
}

func main() {

	intChan := make(chan int, 10)
	boolChan := make(chan bool, 1)

	go wd(intChan)
	go rd(intChan, boolChan)
	for { //这里的话是同时起到阻塞主线程和判断的作用的说O(∩_∩)O哈哈~
		_, ok := <-boolChan
		if !ok {
			break
		}
	}
}
