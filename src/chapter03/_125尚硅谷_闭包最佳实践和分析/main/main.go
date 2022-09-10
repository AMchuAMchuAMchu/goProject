package main

import (
	"fmt"
	"strings"
)

/**
日期:2022/4/28  时间:16:47
@author:冰菓春物咲太实教
*/

func str_buffer(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {

	f01 := str_buffer(".jpg")
	fmt.Println("测试01:", f01("xzxxn"))
	fmt.Println("测试01:", f01("bhzj.jpg"))

}
