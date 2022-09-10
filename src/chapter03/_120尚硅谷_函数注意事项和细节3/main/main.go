package main

import "fmt"

func sum(num1 int, args ...int) int {

	sum := num1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum

}

func main() {
	sum := sum(12, 89, 99, 89, 2236, 4456)
	fmt.Println("结果为:", sum)
}
