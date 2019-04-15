package day09

import "testing"

type ListNod struct {
	Val int
	Next *ListNod
}

func deleteDuplicate (head *ListNod) *ListNod{
	if head != nil && head.Next != nil {
		return head
	}
	temp := head
	for{
		if temp != nil && temp.Next != nil {
			if temp.Next.Val != temp.Val {
				temp = temp.Next
			}else{
				temp.Next = temp.Next.Next
			}
		}else{
			break
		}
	}
	return head
}

func TestDeleteNode(t *testing.T){


	tail := &ListNod{100, nil}
	next := &ListNod{101,tail}
	head := & ListNod{100 ,next}
	duplicate := deleteDuplicate(head)
	t.Log(duplicate.Next.Next)
}