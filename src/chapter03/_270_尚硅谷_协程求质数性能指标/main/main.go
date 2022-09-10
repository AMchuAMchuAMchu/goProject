package main

import (
	"fmt"
	"time"
)

/**
*@BelongsProject: goProject
*@BelongsPackage:
*@Author: 02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
*@CreateTime: 2022-07-24 13:00:57
*@Description: TODO
*@Version: 1.0
 */

func getNum(intChan chan int) {

	for i := 2; i <= 80000; i++ {

		intChan <- i

	}

	close(intChan)

}

func getPrime(intChan chan int, primeChan chan int, boolChan chan bool) {
	for {
		i, ok := <-intChan
		if !ok {
			break
		}
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- i
		}
	}

	boolChan <- true

	fmt.Println("有一个携程执行完毕的书O(∩_∩)O哈哈~...")

}

func main() {

	start := time.Now().UnixMilli()
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 10000)
	boolChan := make(chan bool, 12)

	go getNum(intChan)

	for i := 0; i < 12; i++ {
		go getPrime(intChan, primeChan, boolChan)
	}

	//func() {
	for i := 0; i < 12; i++ {
		<-boolChan
	}
	close(boolChan)
	close(primeChan)

	end := time.Now().UnixMilli()

	fmt.Print("耗时::", end-start, "毫秒,个数:")

	//}()

	//统计个数进行校验...
	//t := 0
	//for  {
	//	_,ok := <- primeChan
	//
	//	if !ok {
	//		break
	//	}
	//	t++
	//}
	//fmt.Print(t)

	//四核(逻辑CPU)==>耗时::552毫秒,个数:7838  耗时::526毫秒,个数:7838  耗时::537毫秒,个数:7838
	//12核(逻辑CPU)==>耗时::298毫秒(刚好是六个物理内核的差距的样子),个数:7838   耗时::292毫秒,个数:7838   耗时::287毫秒,个数:7838
}
