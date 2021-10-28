package _5

import (
	"container/list"
	"testing"
)

//TODO 目前水平不够

//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
func TestReverseKGroup(t *testing.T) {
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
	reverseKGroup(&node, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{}
	//l作为容量为k的栈使用
	l := list.New()

	for {
		//死循环遍历所有元素
		if head != nil {
			//链表插入一个元素
			l.PushFront(head)
			//指向下一个元素
			head = head.Next
			//栈中元素==k,开始弹栈
			if l.Len() == k {
				for i := 0; i < k; i++ {

					value := l.Front().Value.(*ListNode)
					node.Next = value

				}
			}

		} else {
			break
		}
	}
	return node
}
