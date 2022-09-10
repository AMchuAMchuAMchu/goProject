package main

import "fmt"

/**
日期:2022/5/9  时间:17:39
@author:冰菓春物咲太实教
*/

func avg_hens() {
	hen_sum := 0.0
	hen_weight := 0.0
	hens_weight := 0.0
	count := 0
	fmt.Println("请客官输入鸡的个数^_^:")
	fmt.Scan(&hen_sum)
	for i := 0.0; i < hen_sum; i++ {
		count++
		fmt.Println("请输入第", count, "个鸡的重量^_^:")
		fmt.Scan(&hen_weight)
		hens_weight += hen_weight
	}
	fmt.Println("鸡的总体重是:", hens_weight, "千克,鸡的平均体重:", hens_weight/hen_sum, "千克.^_^")
}

func main() {
	//num1 := 0.0
	//fmt.Println("::==>",90.0/num1)

	avg_hens()

}
