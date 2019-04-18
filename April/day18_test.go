package day09

import "testing"

/**
1-100 中的素数
 */
func TestSuShu(t *testing.T){
	for i :=2 ; i < 100; i++ {
		flag := true
		for j :=2 ; j < i; j++ {
			if i % j == 0 {
				flag = false
			}
		}
		if flag{
			//确认是素数
			t.Log(i)
		}
	}
}
