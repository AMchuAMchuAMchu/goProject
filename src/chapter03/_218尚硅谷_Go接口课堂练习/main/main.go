package main

import "fmt"

/**
日期:2022/6/4  时间:7:48
@author:冰菓春物咲太实教
*/

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	//Test03() bool
	Test03()
}

type Stu struct {
}

func (stu Stu) Test01() {

}

func (stu Stu) Test02() {

}

func (stu Stu) Test03() {

}

func main() {
	stu := Stu{}
	var a AInterface = stu
	var b BInterface = stu
	fmt.Println("ok~", a, b)
}
