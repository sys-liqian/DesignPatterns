package a2

import "testing"

//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。

//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
func TestAddTwoNumbers(t *testing.T) {
	node1 := ListNode{
		2,
		&ListNode{
			4,
			&ListNode{
				3, nil}}}
	node2 := ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4, nil}}}

	addTwoNumbers(&node1, &node2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	for l1.Next != nil || l2.Next != nil {

	}
	return nil
}
