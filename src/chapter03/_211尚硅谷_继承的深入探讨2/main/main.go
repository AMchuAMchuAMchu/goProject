package main

import "fmt"

/**
日期:2022/6/3  时间:7:57
@author:冰菓春物咲太实教
*/

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
}

type C struct {
	a A
	//B
}

type Goods struct {
	Name  string
	Price int
}

type Brand struct {
	Who     string
	Address string
}

type PS6 struct {
	Goods
	Brand
}
type PS66 struct {
	*Goods
	*Brand
}

func main() {
	var C C
	//C.A.Name = "雪乃"
	C.a.Name = "777"
	ps6 := PS6{
		Goods: Goods{Name: "m78", Price: 9999},
		Brand: Brand{Who: "索尼", Address: "Japanese"},
	}
	fmt.Println(ps6)
	ps66 := PS66{
		Goods: &Goods{Name: "m78", Price: 9999},
		Brand: &Brand{Who: "索尼", Address: "Japanese"},
	}
	fmt.Println(*ps66.Goods, *ps66.Brand)
}
