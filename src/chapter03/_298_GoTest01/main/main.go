package main

import (
	"fmt"
	"time"
)

/**
*Description==>TODO
*BelongsProject==>goProject
*BelongsPackage==>
*CreateTime==>2022-09-10 17:28:38
*Version==>1.0
*Author==>02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
 */

func main() {

	start := time.Now().UnixMilli()

	channel := make(chan bool, 1)
	//add()
	for i := 0; i < 6; i++ {
		go add(channel)
	}

	for {
		_, err := <-channel
		if !err {
			break
		}
	}

	end := time.Now().UnixMilli()

	fmt.Println(" >> ", end-start, "毫秒") // >>  249 毫秒>>  236 毫秒 >>  253 毫秒

}

/**
 * description==>TODO
 * params==>
 * return==>
 * createTime==>2022/9/10 18:01
 * author==>02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
 */
func add(channel chan bool) {
	tickets := 999999999
	channel <- true
	defer close(channel)
	for {
		tickets--
		if tickets <= 0 {
			return
		}
	}
}
