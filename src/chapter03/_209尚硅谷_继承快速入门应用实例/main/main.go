package main

import "fmt"

/**
日期:2022/6/2  时间:18:36
@author:冰菓春物咲太实教
*/

type Animation struct {
	Name    string
	Author  string
	Time    string
	factory string
}

func (a *Animation) say() {
	fmt.Println("say ok ....")
}

type SDS struct {
	Animation
	Nation string
}

func (S *SDS) say() {
	fmt.Println("SDS..撒野  ")
}
func main() {
	a01 := SDS{
		Animation: Animation{
			"七大罪",
			"铃木央",
			"2011",
			"ANIPLEX",
		},
		Nation: "Japanese",
	}
	a01.Animation.say()
	a01.say()
	fmt.Println(a01)

}
