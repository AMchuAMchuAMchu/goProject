package main

import "fmt"

//
//import "fmt"
//
///**
//日期:2022/4/15  时间:15:49
//@author:冰菓春物咲太实教
//*/
//
//func sum_sub01(num1 int, num2 int) (int,int) {
//	sum := num1 + num2
//	sub := num1 - num2
//	return sum,sub
//}
//
//
//func sum(num1 int, num2 int) int {
//	sum := num1 + num2
//	return sum
//}
//
//func f1(funvar01 func(num1 int ,num2 int) int,num1 int ,num2 int) int {
//	return funvar01(num1 , num2)
//}
//
//type funVar01 func (num1 int ,num2 int) int
//
//func f2(fv funVar01,num1 int ,num2 int) int {
//	return fv(num1, num2)
//}
//
//func sum_sub(num1 int ,num2 int) (sum int , sub int) {
//	sum = num1 + num2
//	sub = num1 - num2
//	return
//}
//
//func main() {
//	calc := sum
//	fmt.Println("计算和差:",calc(99,80))
//	res01 := f1(sum,99,100)
//	fmt.Println("形参:",res01)
//	sum01,sub01 := sum_sub(99,66)
//	fmt.Println("打印和:",sum01,"打印差:",sub01)
//
//}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func sub(num1 int, num2 int) int {
	return num1 - num2
}

func sum1(f1 func(num1 int, num2 int) int, num1 int, num2 int) int {
	return f1(num1, num2)
}

type subType func(num1 int, num2 int) int

func subT(sub subType, num1 int, num2 int) int {
	return sub(num1, num2)
}

func sum_sub(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return

}

func main() {
	s1 := sum(99, 55)
	fmt.Println("1.0打印:", s1)
	fmt.Println("2.0打印:", sum1(sum, 199, 20))
	fmt.Println("3.0打印:", subT(sub, 996, 965))
	sum, sub := sum_sub(520, 419)
	fmt.Println("4.0打印和:", sum, "4.0打印差:", sub)
}
