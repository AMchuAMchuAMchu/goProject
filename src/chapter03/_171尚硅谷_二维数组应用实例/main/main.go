package main

import "fmt"

/**
日期:2022/5/16  时间:11:41
@author:冰菓春物咲太实教
*/

func main() {

	classs := [3][5]int{{}, {}, {}}
	sum := 0
	sum01 := 0
	sum02 := 0
	sum03 := 0
	person := 0
	for i := 0; i < len(classs); i++ {
		for j := 0; j < len(classs[i]); j++ {
			fmt.Println("请客官输入第", i+1, "个班第", j+1, "个同学的分数^_^:")
			fmt.Scan(&classs[i][j])
			sum += classs[i][j]
			person += 1
		}
	}

	for i := 0; i < len(classs[0]); i++ {
		sum01 += classs[0][i]
		sum02 += classs[1][i]
		sum03 += classs[2][i]
	}

	fmt.Println("久等啦~~^_^所有班级的平均分是:", sum/person)
	//fmt.Println(person)
	fmt.Println("第一个班级的平均分是(≧∇≦)ﾉ:", sum01/(person/3))
	//fmt.Println(sum01)
	fmt.Println("第二个班级的平均分是(≧∇≦)ﾉ:", sum02/(person/3))
	fmt.Println("第三个班级的平均分是(≧∇≦)ﾉ:", sum03/(person/3))

}
