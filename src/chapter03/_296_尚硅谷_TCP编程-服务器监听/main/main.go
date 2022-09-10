package main

import (
	"fmt"
	"net"
)

/**
*@BelongsProject: goProject
*@BelongsPackage:
*@Author: 02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
*@CreateTime: 2022-09-07 14:07:48
*@Description: TODO
*@Version: 1.0
 */

func main() {

	fmt.Println("go 版本 web 开始干活的说 ^_^...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	defer listen.Close()

	if err != nil {
		fmt.Println("完蛋,连接不上...>_<", err)
		return
	} else {
		fmt.Println("呀哈喽~~", listen.Addr().String())
	}

	for {
		fmt.Println("侦听中...")
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("果咩纳塞...>_<连接不上的说...", err)
		} else {
			fmt.Println("夏岛悬疑...", conn.RemoteAddr().String(), "已成功的连接上来的说^_^...")
			process(conn)
		}

	}

}

func process(conn net.Conn) {

	defer conn.Close()

	for {

		buf := make([]byte, 1024)

		fmt.Println("服务器在等待xxx信息输入的说^_^...")

		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn.Read Err ==> ", err)
		}

		fmt.Println("服务器读取到客户端::", conn.RemoteAddr().String(), "的信息是==>", string(buf))

	}

}
