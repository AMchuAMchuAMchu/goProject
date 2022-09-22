package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/**
*Description==>TODO
*BelongsProject==>goProject
*BelongsPackage==>
*CreateTime==>2022-09-22 09:05:26
*Version==>1.0
*Author==>02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
 */

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("...残念~~")
		return
	}

	fmt.Println(">>", conn)

}
