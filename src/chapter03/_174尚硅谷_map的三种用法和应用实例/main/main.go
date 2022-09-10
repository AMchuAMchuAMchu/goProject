package main

import "fmt"

/**
日期:2022/5/18  时间:20:09
@author:冰菓春物咲太实教
*/

func main() {

	studentList := make(map[string]map[string]string)
	studentList["第一个学生"] = map[string]string{"name": "雪乃", "sex": "女"}
	studentList["第二个学生"] = map[string]string{"name": "宇佐美", "sex": "女"}
	studentList["第三个学生"] = map[string]string{"name": "伊万莉玛利亚", "sex": "女"}
	fmt.Println(":==>", studentList)
	for s, m := range studentList {
		fmt.Println(s)
		fmt.Println(m)
		for s2, s3 := range m {
			fmt.Println(s2)
			fmt.Println(s3)
		}
	}

}
