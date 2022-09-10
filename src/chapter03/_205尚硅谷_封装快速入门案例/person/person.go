package person

import "fmt"

/**
日期:2022/6/1  时间:21:51
@author:冰菓春物咲太实教
*/

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name:   name,
		age:    0,
		salary: 0,
	}
}

func (p *person) SetAge(age int) {
	if !(age > 0 && age < 150) {
		fmt.Println("输入不合法!!!(•́へ•́╬)")
		return
	}
	p.age = age
	fmt.Println("修改成功!!^_^")
}

func (p *person) GetAge() int {
	return p.age
}
