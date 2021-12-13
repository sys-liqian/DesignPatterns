package algorithm

import (
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
	PreOrder(&node)
	fmt.Println()
	InfixOrder(&node)
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
