package main

import "fmt"

/**
日期:2022/5/8  时间:20:18
@author:冰菓春物咲太实教
*/

//func main() {
//	for i := 0; i < 10; i++ {
//
//	var userIput string
//	fmt.Println("请客官输入:")
//	fmt.Scan(&userIput)
//	switch userIput {
//	case "1":
//		fmt.Println("这是一月,一共有31天...")
//	case "2":
//		fmt.Println("这是一月,一共有29天...")
//	case "3":
//		fmt.Println("这是一月,一共有31天...")
//	case "4":
//		fmt.Println("这是一月,一共有30天...")
//	case "5":
//		fmt.Println("这是一月,一共有31天...")
//	case "6":
//		fmt.Println("这是一月,一共有30天...")
//	case "7":
//		fmt.Println("这是一月,一共有31天...")
//	case "8":
//		fmt.Println("这是一月,一共有31天...")
//	case "9":
//		fmt.Println("这是一月,一共有30天...")
//	case "10":
//		fmt.Println("这是一月,一共有31天...")
//	case "11":
//		fmt.Println("这是一月,一共有30天...")
//	case "12":
//		fmt.Println("这是一月,一共有31天...")
//	}
//	}
//
//
//
//
//
//
//
//}

//func main() {
//	var count int = 1
//
//	rand.Seed(time.Now().UnixNano())
//	var random int = rand.Intn(100)
//	for i := 0; i < 10; i++ {
//
//		var user_input int
//
//		fmt.Println("第", count, "次请客官输入:")
//		fmt.Scan(&user_input)
//
//		if user_input > random {
//			fmt.Println("偏大了...")
//		} else if user_input < random {
//			fmt.Println("偏小了...")
//
//		}
//
//		if random == user_input {
//			if count == 1 {
//
//				fmt.Println("你真是个天才!!!")
//				break
//			} else if count >= 2 && count <= 3 {
//				fmt.Println("你很聪明,赶上我了..")
//				break
//			} else if count >= 4 && count <= 9 {
//				fmt.Println("一般般...")
//				break
//			} else if count == 10 {
//				fmt.Println("可算猜对了...")
//				break
//			} else {
//				fmt.Println("该说你点啥好呢....")
//				break
//			}
//
//		}
//		count++
//
//	}
//
//}

//func main() {
//
//	var count int = 0
//	for i := 2; i <= 1000; i++ {
//		var  flag bool = true
//		for j := 2; float64(j) <= math.Sqrt(float64(i)); j++ {
//			if i%j == 0 {
//				flag = false
//				break//168
//			}
//		}
//
//		if flag {
//			count++
//			fmt.Println("第",count,"个质数:",i)
//		}
//	}
//
//
//
//
//
//}

func main() {

	var letter int = 97
	var b_letter byte = 65
	count := 0
	for i := 0; i < 26; i++ {
		count++
		fmt.Println("第", count, "个小写:", string(letter))
		fmt.Println("第", count, "个大写:", string(b_letter))
		letter += 1
		b_letter += 1
	}

}
