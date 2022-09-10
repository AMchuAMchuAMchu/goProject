package main

import (
	"fmt"
	"sort"
)

/**
日期:2022/6/18  时间:21:57
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

type xinei struct {
	Age  int
	Name string
}

type xinei01 []xinei

func (xn xinei01) Len() int {
	return len(xn)
}

func (xn xinei01) Swap(i, j int) {
	temp := xn[i]
	xn[i] = xn[j]
	xn[j] = temp
}

func (xn xinei01) Less(i, j int) bool {
	return xn[i].Age > xn[j].Age
}

func main() {

	xn := xinei01{
		{
			Age:  15,
			Name: "雪乃",
		},
		{
			Age:  17,
			Name: "樱满集",
		},
		{
			Age:  27,
			Name: "布兰德",
		},
		{
			Age:  16,
			Name: "楪祈",
		},
		{
			Age:  23,
			Name: "艾斯德斯",
		},
	}

	sort.Sort(xn)
	fmt.Println(xn)

}
