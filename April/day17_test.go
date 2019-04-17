package day09

import "testing"

func TestMergeSort(t *testing.T){
	arr := []int{44,33,22,55,66,11}
	mergeSort(arr,0,len(arr)-1)
}

func mergeSort(arr []int , start , end  int){
	if start >= end{
		return
	}
	mid := (start + end) / 2
	mergeSort(arr ,start , mid)
	mergeSort(arr, mid+1 , end)
	merge (arr ,start , mid ,end)
}

func merge(arr []int ,start ,mid ,end int){
	i := start
	j := end
	k := 0
	tmpArr := make([]int , j-1+1)
	for ;i <= mid && j <= end;k++{
		if arr [i] <= arr[j] {
			arr[k] = arr[i]
			i++
		}else{
			arr[k] = arr[j]
			j++
		}
	}
	for ;i <= mid; i++ {
		arr [k] = arr [i]
		i++
	}

	for ;j <= mid; j++ {
		arr [k] = arr [j]
		j++
	}
	copy(arr[start:end+1],tmpArr)

}


