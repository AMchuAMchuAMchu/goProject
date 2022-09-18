package _299_benchmarkTest01

import (
	"fmt"
	"time"
)

/**
*Description==>TODO
*BelongsProject==>goProject
*BelongsPackage==>
*CreateTime==>2022-09-18 11:13:11
*Version==>1.0
*Author==>02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
 */

func main() {

	start := time.Now().UnixMilli()
	for i := 1; i < 999999999; i++ {

	}
	end := time.Now().UnixMilli()

	fmt.Println("耗时:", end-start, "毫秒...") //

}
