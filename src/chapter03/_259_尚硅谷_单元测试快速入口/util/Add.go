package util

/**
日期:2022/7/4  时间:20:51
@author:02雪乃赤瞳楪祈校条祭制作委员会
*/

func Add(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
