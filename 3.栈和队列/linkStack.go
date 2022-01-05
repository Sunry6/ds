package main

import (
	"fmt"
	"sync"
)

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

// 入栈

func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		// 否则新元素插入链表的头部
		// 原来的链表
		preNode := stack.root

		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 原来的链表链接到新元素后面
		newNode.Next = preNode

		// 将新节点放在头部
		stack.root = newNode
	}

	// 栈中元素数量+1
	stack.size = stack.size + 1
}

// 出栈

func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	topNode := stack.root
	v := topNode.Value

	stack.root = topNode.Next
	stack.size -= 1

	return v
}

// 获取栈顶元素

func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}

	v := stack.root.Value
	return v
}

// 栈大小

func (stack *LinkStack) Size() int {
	return stack.Size()
}

// 栈是否为空

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Push("cat")
	fmt.Println("size:", linkStack.Size())
}
