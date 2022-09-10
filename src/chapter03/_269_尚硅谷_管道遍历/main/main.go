package main

import (
	"fmt"
	"time"
)

func wd(intChan chan int) {
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("写入:", i)
		time.Sleep(time.Microsecond * 10)
	}
	close(intChan)
}

func rd(intChan chan int) {
	for v := range intChan {
		fmt.Println("读取:", v)
		time.Sleep(time.Microsecond * 10)
	}
}

func main() {

	intChan := make(chan int, 10)

	go wd(intChan)
	go rd(intChan)

	time.Sleep(time.Second) //不要出现协程还没结束主线程就已经结束了的情况...
}
