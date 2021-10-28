package a206

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
func TestReverseList(t *testing.T) {
	node := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{5, nil}}}},
	}
	_ = reverseList(&node)
}

/*
p0 -> p1 -> p2 -> p3 ->p4

//缓存当前节点的下一个节点
tmp = p1

prev <- p0 -> p1 -> p2 -> p3 ->p4

*/

func reverseList(head *ListNode) *ListNode {
	//当前指针前一个指针
	var ptr1 *ListNode
	//当前指针
	ptr2 := head
	for {
		//当前指针不为空一直循环
		if ptr2 != nil {
			//tmp 为临时指针，指向当前指针的下一个地址
			tmp := ptr2.Next
			//当前指针指向前一个指针
			ptr2.Next = ptr1
			//当前指针前一个指针后移到当前指针处
			ptr1 = ptr2
			//当前指针后移
			ptr2 = tmp
		} else {
			break
		}
	}
	return ptr1
}
