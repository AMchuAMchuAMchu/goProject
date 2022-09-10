package main

import "fmt"

func sum(n1, n2 float32) float32 {
	fmt.Printf("n1 type = %T\n", n1)
	return n1 + n2
}

func swap(n1 int, n2 int) (int, int) {

	n3 := n1
	n1 = n2
	n2 = n3
	return n1, n2

}

func main() {
	//sum := sum(1,2)
	//fmt.Println("结果为:",sum)
	//main()
	n1, n2 := swap(10, 20)
	fmt.Println("n1 = ", n1, "n2 = ", n2)

}
