package model

/**
日期:2022/5/27  时间:18:00
@author:冰菓春物咲太实教
*/
type animal struct {
	Name string
	Age  int
}

func GetAnimalInstance(Name string, Age int) animal {
	return animal{
		Name: Name,
		Age:  Age,
	}
}
