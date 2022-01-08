package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

type LinkQueue2 struct {
	root *LinkNode3
	size int
	lock sync.Mutex
}

type LinkNode3 struct {
	Next  *LinkNode3
	Value *TreeNode
}

func (queue *LinkQueue2) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode3)
		queue.root.Value = v
	} else {
		newNode := new(LinkNode3)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		nowNode.Next = newNode
	}
	queue.size += 1
}

// 出队

func (queue *LinkQueue2) Remove() TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	topNode := queue.root
	queue.root = topNode.Next
	queue.size -= 1

	return TreeNode{}
}

// PreOrder 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data, " ")
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// MidOrder 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	fmt.Println(tree.Data, " ")
	MidOrder(tree.Right)
}

// PostOrder 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Println(tree.Data, " ")
}

// LayerOrder 层次遍历
func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}
	queue := new(LinkQueue2)
	queue.Add(treeNode)
	for queue.size > 0 {
		element := queue.Remove()
		fmt.Println(element.Data)
		if element.Left != nil {
			queue.Add(element.Left)
		}
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

func main() {

}
