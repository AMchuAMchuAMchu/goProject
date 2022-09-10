package main

import "fmt"

/**
日期:2022/5/18  时间:20:52
@author:冰菓春物咲太实教
*/

func main() {

	//map01 := make(map[string]map[string]string)
	//map01["num01"] = map[string]string{"name":"雪乃","sex":"女"}
	//map01["num02"] = map[string]string{"name":"由比滨结衣","sex":"女"}
	//map01["num03"] = map[string]string{"name":"宇佐美","sex":"女"}
	//fmt.Println(map01)

	animation := make([]map[string]string, 2)
	animation[0] = make(map[string]string)
	animation[1] = make(map[string]string)
	animation[0]["name"] = "雪之下雪乃"
	animation[0]["sex"] = "雪"
	animation[1]["name"] = "由比滨结衣"
	animation[1]["sex"] = "衣"
	for i, m := range animation {
		fmt.Println(i, m)
	}

}
