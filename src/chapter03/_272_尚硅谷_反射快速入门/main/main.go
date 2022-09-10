package main

import (
	"fmt"
	"reflect"
)

/**
*@BelongsProject: goProject
*@BelongsPackage:
*@Author: 02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
*@CreateTime: 2022-07-25 12:50:41
*@Description: TODO
*@Version: 1.0
 */

func testReflect(item interface{}) {

	rT := reflect.TypeOf(item)
	rV := reflect.ValueOf(item)
	sum := 11 + rV.Int()
	fmt.Printf("类型:%v;值:%v;输出运算结果:%v;RV结果的类型:%T", rT, rV, sum, rV)
}

func main() {
	num := 100
	testReflect(num)
}
