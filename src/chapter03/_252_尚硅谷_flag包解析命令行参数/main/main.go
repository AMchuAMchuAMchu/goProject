package main

import (
	"flag"
	"fmt"
)

/**
日期:2022/6/28  时间:15:03
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func main() {

	var user string
	var password string
	var host string
	var port int
	//var xx int
	//xx := 9
	//var xx = 9
	//fmt.Println(xx)

	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&password, "p", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机默认为localhost")
	flag.IntVar(&port, "P", 02, "端口号默认为02")

	flag.Parse()

	fmt.Printf(""+
		"user::%v\n"+
		"password::%v\n"+
		"host::%v\n"+
		"port::%v\n"+
		"", user, password, host, port)

}
