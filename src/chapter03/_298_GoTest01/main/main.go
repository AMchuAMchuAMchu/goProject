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

	go add()

	end := time.Now().UnixMilli()

	fmt.Println(" >> ", end-start, "毫秒") // >>  249 毫秒>>  236 毫秒 >>  253 毫秒

}

func add() {
	tickets := 999999999
	for {
		tickets--
		if tickets <= 0 {
			return
		}
	}
}
