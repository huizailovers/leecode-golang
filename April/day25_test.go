package day09

import (
	"fmt"
	"testing"
)

func TestHuiWen(t *testing.T){
	wen := huiWen()
	fmt.Println(wen)
}

func huiWen() bool{
	str := "A man, a plan, a canal: Panama"
	i := 0
	j := len(str)-1
	flag := true
	for i< j {
		//忽略空格 忽略大小写
		if str[i]  == str[j] {
			i ++
			j --
		}else{
			flag = false
			break
		}
	}
	return flag
}