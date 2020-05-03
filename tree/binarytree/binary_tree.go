package binarytree

import "fmt"

type TreeNode struct {
	Data  string    // 存放当前node数据
	Left  *TreeNode // 记录左子树node信息
	Right *TreeNode // 记录右子树node信息
}

// 先序遍历(根节点 -> 左子树 -> 右子树)
func PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%s ", node.Data)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func MidOrder(node *TreeNode) {
	if node == nil {
		return
	}
	MidOrder(node.Left)
	fmt.Printf("%s ", node.Data)
	MidOrder(node.Right)
}

func PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Printf("%s ", node.Data)
}
