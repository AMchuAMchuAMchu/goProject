package main

import "fmt"

/**
日期:2022/6/3  时间:8:30
@author:冰菓春物咲太实教
*/

type Usb interface {
	start()
	stop()
}

type Phone struct {
}

type Camera struct {
}

func (p Phone) start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) stop() {
	fmt.Println("手机停止工作...")
}

func (c Camera) start() {
	fmt.Println("摄像机开始工作...")
}

func (c Camera) stop() {
	fmt.Println("摄像机停止工作...")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {

	var p Phone
	var c Camera
	var computer Computer
	computer.Working(p)
	computer.Working(c)

}
