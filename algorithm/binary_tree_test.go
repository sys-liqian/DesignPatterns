package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

/*
			A
		B       E
	  C		   F   G
        D         I H

*/

type Node struct {
	V string
	L *Node
	R *Node
}

//二叉树的遍历
func TestTraverse(t *testing.T) {
	node := Node{
		V: "A",
		L: &Node{
			V: "B",
			L: &Node{
				V: "C",
				R: &Node{V: "D"},
			},
		},
		R: &Node{
			V: "E",
			L: &Node{V: "F"},
			R: &Node{
				V: "G",
				L: &Node{V: "I"},
				R: &Node{V: "H"},
			},
		},
	}
	//PreOrder(&node)
	//fmt.Println()
	PreOrder1(&node)
	//InfixOrder(&node)
	fmt.Println()
}

// 根->左->右
// 前序遍历,根节点在左右子树之前
func PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("%s,", node.V)
		PreOrder(node.L)
		PreOrder(node.R)
	}
}

/*
 输出A ,A入栈
 输出B ,B入栈
 输出C ,C入栈, C左孩子为空,所以 node=node.L node=nil,进入第二个判断
 取栈顶元素为C,node = node.R ,然后C出栈
 输出D ,D入栈, D左孩子为空, node = nil,进入第二个判断
 取栈顶元素为D,node = node.R ,D出栈, node=nil, B出栈,node=nil,A出栈 node= E

*/
func PreOrder1(node *Node) {
	stack := list.New()
	for node != nil || stack.Len() != 0 {
		if node != nil {
			fmt.Printf("%s", node.V)
			stack.PushFront(node)
			node = node.L
		}
		if node == nil && stack.Len() != 0 {
			back := stack.Front()
			n := (back.Value).(*Node)
			node = n.R
			stack.Remove(back)
		}
	}
}

// 左->根->右
// 中序遍历,根节点在左右子树之间
func InfixOrder(node *Node) {
	if node != nil {
		InfixOrder(node.L)
		fmt.Printf("%s,", node.V)
		InfixOrder(node.R)
	}
}

// 左->右->根
// 后序遍历,根节点在左右子树之后
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.L)
		PostOrder(node.R)
		fmt.Printf("%s,", node.V)
	}
}
