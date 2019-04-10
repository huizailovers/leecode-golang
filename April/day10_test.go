package day09

import "testing"

/**
删除数组中的重复项 返回数组的长度
 */
func removeDuplicates(nums []int) int{

	if len(nums) == 0{
		return 0
	}

	i := 0

	for j :=1; j<len(nums);j++{
		if nums[i] != nums[j]{
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}


func TestRemoveDuplicates(t *testing.T){
	nums := []int{1,1,2,3}
	len := removeDuplicates(nums)
	t.Log("len =" , len)
}