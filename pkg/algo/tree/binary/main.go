package main

import (
	"algo/pkg/algo/tree/binary/binarytree"
	"fmt"
)

/*
                .-----.
               /       \
              (    A    )
            /  `.     ,'  \
           /     `---'     \
          /                 \
         /                   \
      .-----.             .-----.
     /       \           /       \
    (    B    )         (    C    )
     `/     ,\           `/     ,'
     / `---'  \          / `---'
    /          \        /
   /            \      /
 .---.        .---.  .---.
(  D  )      (  E  )(  F  )
 `---'        `---'  `---'

		先序遍历
			A B D E C F
		中序遍历
			D B E A F C
		后序遍历
			D E B F C A
		层次遍历
			A B C D E F
*/
func main() {
	// 构造一棵树
	tree := &binarytree.TreeNode{Data: "A"}
	tree.Left = &binarytree.TreeNode{Data: "B"}
	tree.Right = &binarytree.TreeNode{Data: "C"}

	tree.Left.Left = &binarytree.TreeNode{Data: "D"}
	tree.Left.Right = &binarytree.TreeNode{Data: "E"}

	tree.Right.Left = &binarytree.TreeNode{Data: "F"}

	// 二叉树先序遍历
	fmt.Println("先序遍历")
	binarytree.PreOrder(tree)
	fmt.Println()
	// 二叉树中序遍历
	fmt.Println("中序遍历")
	binarytree.MidOrder(tree)
	fmt.Println()
	//	二叉树后序遍历
	fmt.Println("后序遍历")
	binarytree.PostOrder(tree)
	fmt.Println()

	// 二叉树层次遍历
	fmt.Println("层次遍历")
	binarytree.BFS(tree)
	fmt.Println()
}
