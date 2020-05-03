package binarytree

import (
	"fmt"
	"sync"
)

func BFS(node *TreeNode) {
	if node == nil {
		return
	}
	// 新建队列
	queue := new(LinkQueue)

	// 根节点入队
	queue.Add(node)
	// 进行队列非空校验，不为空进行遍历
	for queue.size > 0 {
		// 从顶部开始出队元素
		currNode := queue.Remove()
		fmt.Printf("%s ", currNode.Data)
		// 判断是否存在左子树，存在则进行入队
		if currNode.Left != nil {
			queue.Add(currNode.Left)
		}
		// 判断是否存在右子树，存在则进行入队
		if currNode.Right != nil {
			queue.Add(currNode.Right)
		}
	}
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

// 链表队列
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

// 入队
func (queue *LinkQueue) Add(node *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果栈顶为空，则新建一个节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = node
	} else {
		// 不为空，则在链表尾部增加一个节点
		// 新建节点
		newNode := new(LinkNode)
		newNode.Value = node

		// 遍历链表，直至尾部
		currNode := queue.root
		for currNode.Next != nil {
			currNode = currNode.Next
		}

		// 当前为链表尾部，将新建节点插入尾部
		currNode.Next = newNode
	}
	// 队列大小进行加1
	queue.size += 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 判断当前队列是否为nil
	if queue.size <= 0 {
		panic("queue size over limit")
	}
	// 顶部元素出队
	topNode := queue.root
	currNode := topNode.Value

	// 将顶部元素的后继链接链表上
	queue.root = topNode.Next
	// 将队列大小进行减1
	queue.size -= 1
	// 返回出队Node元素
	return currNode
}

// 队列元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}
