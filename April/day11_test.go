package day09

import "testing"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode{
	temp := head
	var pre = &ListNode{}
	for temp != nil{
		second := temp.Next
		temp.Next = pre
		pre = temp
		temp = second
	}
	return pre
}

func TestReverseList(t *testing.T){
	second := &ListNode{2,nil}
	first := &ListNode{1,second}
	list := reverseList(first)
	t.Log(list.Val)
}

