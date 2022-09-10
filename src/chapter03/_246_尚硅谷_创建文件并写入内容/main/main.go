package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
日期:2022/6/27  时间:17:45
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/
func main() {
	fileStr := "D:\\seldom\\rd\\Go_ProjectAll\\goProject\\src\\chapter03\\assets\\image\\_03ysn.txt\n"
	data, err := os.OpenFile(fileStr, os.O_WRONLY|os.O_CREATE, 006)
	defer data.Close()
	if err != nil {
		fmt.Println("!!浸警告!!(•́へ•́╬)可能出BUG啦...")
	}
	fileObj := bufio.NewWriter(data)
	strAdd := "\n结城明日奈，日本轻小说《刀剑神域》及其衍生作品中的女主角。 [1-2]  因好奇借用哥哥\n" +
		"结城浩一郎的NERvGear进入SAO的网络游戏新手，并用真名取游戏角色名字。卷入了死亡游戏后内心一\n" +
		"直恐惧而感到害怕，一直想尽快突破100层回到现实世界，被攻略组称为顶级公会的攻略之鬼，因为和桐人\n" +
		"不和而决斗，后开始注意起桐人，直到认识“活在”SAO里的桐谷和人后改变想法，也因此喜欢上桐谷和人\n" +
		"并在游戏里结婚。\n"
	for i := 0; i < 10; i++ {
		line, err := fileObj.WriteString(strAdd)
		if err != nil {
			fmt.Println("结束...了😭😭😭😲😲😲")
			return
		}
		fmt.Println("==>", line)
	}
	fileObj.Flush()
}
