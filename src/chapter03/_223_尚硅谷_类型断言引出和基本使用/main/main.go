package main

import "fmt"

/**
日期:2022/6/19  时间:19:34
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

type Point struct {
	x int
	y int
}

func main() {

	//var a interface{}
	//var p Point  = Point{
	//	x: 0,
	//	y: 0,
	//}
	//a = p
	//var p1 Point
	//fmt.Println(p.x,p1.x)
	//p1 = a.(Point)
	//if _,b2 := a.(Point);b2{
	//	fmt.Println("00000")
	//}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		switch i {
		case 50:
			fmt.Println("==>", i)
			return
		}
	}

}
