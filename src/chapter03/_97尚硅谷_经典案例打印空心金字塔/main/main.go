package main

import "fmt"

/**
日期:2022/4/5  时间:11:48
@author:冰菓春物咲太实教
*/

func main() {

	r_length := 0
	fmt.Println("请输入要打印的星星的行数O(∩_∩)O哈哈~:")
	fmt.Scanln(&r_length)

	for r := 1; r <= r_length; r++ {

		for k := 0; k <= r_length-r; k++ {
			fmt.Print(" ")
		}

		for c := 1; c <= 2*r-1; c++ {
			if c == 1 || c == 2*r-1 || r == r_length {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}

		fmt.Println()

	}

	/*
		上半部分:
		  *      星星数目:1    规律一:2 * 当前行数 - 1     星星左边的空格数目:2   规律二:总行数 - 当前行数
		 ***     星星数目:3    规律一:2 * 当前行数 - 1     星星左边的空格数目:1   规律二:总行数 - 当前行数
		*****    星星数目:5    规律一:2 * 当前行数 - 1     星星左边的空格数目:0   规律二:总行数 - 当前行数

		下半部分:
		*****    星星数目:5
		 ***     星星数目:3
		  *      星星数目:1

	*/

	//r_length := 0
	//fmt.Println("请输入需要打印的菱形的行数Thanks♪(･ω･)ﾉ:")
	//fmt.Scanln(&r_length)
	//for r := 1;r <= r_length;r++ {
	//
	//	for k := 1; k <= r_length - r;k++ {
	//		fmt.Print(" ")
	//	}
	//
	//
	//	for c := 1;c <= (2 * r - 1); c++ {
	//		if c == 1 || c == (2 * r -1) || r == r_length {
	//			fmt.Print("*")
	//		}else {
	//			fmt.Print(" ")
	//		}
	//	}
	//
	//
	//
	//	fmt.Println()
	//
	//}

	/*
		  *      星星数:1  规律:2 * 行数 - 1    空格的个数:2   规律:总行数 - 当前行数
		 ***     星星数:2  规律:2 * 行数 - 1    空格的个数:1   规律:总行数 - 当前行数
		*****    星星数:3  规律:2 * 行数 - 1    空格的个数:0   规律:总行数 - 当前行数


	*/
	//row_length := 0
	//fmt.Println("请输入打印的金字塔的行数(≧∇≦)ﾉ:")
	//fmt.Scanln(&row_length)
	//for r := 1;r <= row_length; r++ {
	//
	//	for k := 1;k <= row_length - r;k++ {
	//		fmt.Print(" ")
	//	}
	//
	//	for c := 1;c <= 2 * r - 1;c++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//
	//}

	//for r := 1;r <= 5;r++ {
	//	for c := 1;c <= r;c ++ {
	//		fmt.Print(" 😂 ")
	//	}
	//	fmt.Println()
	//}

	//for r := 5;r >= 1;r-- {
	//	for c := 1;c <= r;c ++ {
	//		fmt.Print(" 😂 ")
	//	}
	//	fmt.Println()
	//}

	//
	//row_length := 0//接收用户输入的打印的行数
	//column_length := 0//接收用户输入的打印的列数
	//fmt.Println("请输入打印的行数(≧∇≦)ﾉ:")
	//fmt.Scanln(&row_length)
	//fmt.Println("请输入打印的列数(≧∇≦)ﾉ:")
	//fmt.Scanln(&column_length)
	//for r := 1;r <= row_length;r++ {
	//	for c := 1;c < column_length;c++{
	//		fmt.Print(" 😎 ")
	//	}
	//	fmt.Println()
	//}

}
