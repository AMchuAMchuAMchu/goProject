package util

import (
	"fmt"
	"testing"
)

/**
日期:2022/7/4  时间:20:54
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func TestAdd(t *testing.T) {
	sum := Add(100)
	if sum != 5050 {
		fmt.Println("!!警告!!(•́へ•́╬)")
		t.Fatalf("!!警告!!(•́へ•́╬)")
	} else {
		fmt.Println("哦咩爹多~~^_^")
		t.Logf("哦咩爹多~~^_^")
	}
}
