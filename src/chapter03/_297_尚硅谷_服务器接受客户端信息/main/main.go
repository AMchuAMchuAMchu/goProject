package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
*Description==>TODO
*BelongsProject==>goProject
*BelongsPackage==>
*CreateTime==>2022-09-08 19:35:51
*Version==>1.0
*Author==>02雪乃赤瞳楪祈校条祭制作委员会 wyq_start
 */

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("完蛋..残念~~")
		return
	} else {
		fmt.Println(" >> ", conn.RemoteAddr().String())
	}

	for {

		bufioStr := bufio.NewReader(os.Stdin)

		message01, err := bufioStr.ReadString('\n')

		if err != nil {
			fmt.Println("bufioStr.ReadString Err ==> ", err)
		}

		message01 = strings.Trim(message01, "\n")

		if message01 == "exit" {
			fmt.Println("bye bye ...")
			return
		}

		writeLength, err := conn.Write([]byte(message01))

		if err != nil {
			fmt.Println("conn.Write Err ==>", writeLength)
			return
		}

		fmt.Println("客户端::", conn.RemoteAddr().String(), "发送信息差长度==>", writeLength)

	}

}
