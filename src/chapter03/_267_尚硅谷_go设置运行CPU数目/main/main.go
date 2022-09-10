package main

import (
	"fmt"
	"runtime"
)

//
//import (
//	"fmt"
//	"runtime"
//)
//
///**
//日期:2022/7/7  时间:19:08
//@author:02雪乃赤瞳楪祈校条祭制作委员会
//*/
//
//
//
//
func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum++:", cpuNum)
	runtime.GOMAXPROCS(cpuNum - 2)
}

//
