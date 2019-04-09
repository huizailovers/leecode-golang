package day09_test

import "testing"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoList(l1 *ListNode , l2 *ListNode) *ListNode{
	temp := & ListNode{}
	head := temp

	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val {
			temp.Next = l1
			temp = temp.Next
			l1 = l1.Next
		}else{
			temp.Next = l2
			temp = temp.Next
			l2 = l2.Next
		}

	}
	if l1 != nil{
		temp.Next = l1
	}
	if l2 != nil{
		temp.Next = l1
	}

	return head.Next
}

func  TestMerge(t *testing.T){
	l1 := &ListNode{1,nil}
	l2 := &ListNode{2,nil}
	list := mergeTwoList(l1, l2)
	t.Log(list.Val)
}