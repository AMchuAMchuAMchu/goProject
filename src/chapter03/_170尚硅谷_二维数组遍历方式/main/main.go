package main

import "fmt"

/**
日期:2022/5/16  时间:11:03
@author:冰菓春物咲太实教
*/

func main() {
	arr01_01 := [4]string{"约定的梦幻岛", "鬼灭之刃", "我们无法一起起学习", "石纪元"}
	arr01_02 := [4]string{"樱满集", "蝶祈"}
	arr01_03 := [4]string{"未来日记", "我妻由乃", "天野雪辉", "宙斯"}
	arr01_04 := [4]string{"火影忍者", "鸣人", "佐助"}
	arr01 := [4][4]string{arr01_01, arr01_02, arr01_03, arr01_04}
	for i := 0; i < len(arr01); i++ {
		for j := 0; j < len(arr01[i]); j++ {
			fmt.Print("第", i+1, "行第", j+1, "个值^_^:", arr01[i][j], " || ")
		}
		fmt.Println()
	}
	fmt.Println("==>")
	for i, strings := range arr01 {
		for i2, s := range strings {
			fmt.Print("第", i+1, "行第", i2+1, "个值:", s, "##")
		}
		fmt.Println()
	}

}
