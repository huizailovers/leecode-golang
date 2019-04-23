package day09

import "testing"

func TestReverseArray(t *testing.T){
   arr := []int{11,22,33,44}
   ReverseArray(arr)
   for i :=0; i< len(arr);i++{
   	t.Log(arr[i])
   }
}

func ReverseArray(arr []int){
	i :=0
	for j:= len(arr)-1; j >=0;j--{
		arr[i] ,arr[j] = arr[j] ,arr[i]
		i++
	}

}
