package main

import (
	"encoding/json"
	"fmt"
)

/**
日期:2022/5/22  时间:15:09
@author:冰菓春物咲太实教
*/

type Monster struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Hobby string `json:"hobby"`
}

func main() {

	m01 := Monster{
		Name:  "zero two",
		Age:   "17",
		Hobby: "Darling",
	}
	fmt.Println(m01)
	v, _ := json.Marshal(m01)

	fmt.Println(v)
	fmt.Println(string(v))

}
