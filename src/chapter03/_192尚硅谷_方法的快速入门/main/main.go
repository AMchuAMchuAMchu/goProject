package main

import "fmt"

/**
日期:2022/5/22  时间:15:35
@author:冰菓春物咲太实教
*/

type Person struct {
	Name  string
	Hobby string
}

func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人...(≧∇≦)ﾉ")
}

func (p Person) jisuan() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("计算1-1000结果:==>", sum)
}

func (p Person) jisuan2(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("计算1-n结果:==>", sum)
}

func (p Person) sum(n1 int, n2 int) int {
	return (n1 + n2)
}

func main() {
	p := Person{
		Name:  "Jerry",
		Hobby: "看番...",
	}
	res01 := p.sum(99, 999)
	p.jisuan()
	p.jisuan2(100)
	p.speak()
	fmt.Println(res01)
}
