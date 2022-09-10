package main

import "fmt"

func sum(n1 int, n2 int) int {

	sum := n1 + n2
	return sum

}

var (
	fun01 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	s := sum(88, 123)
	fmt.Println("匿名函数一:", s)
	res01 := func(n1 int, n2 int) int {
		return n1 - n2
	}(12, 34)
	fmt.Println("匿名函数二:", res01)
	fmt.Println("匿名函数三:", fun01(1, 1))

}
