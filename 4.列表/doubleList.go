package main

import "sync"

type DoubleList struct {
	head *ListNode
	tail *ListNode
	len  int
	lock sync.Mutex
}

type ListNode struct {
	pre   *ListNode
	next  *ListNode
	value string
}

// GetValue 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// GetPre 获取节点的前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// GetNext 获取节点的后继节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// HashNext 是否存在后驱节点
func (node *ListNode) HashNext() bool {
	return node.next != nil
}

// HashPre 是否存在前驱节点
func (node *ListNode) HashPre() bool {
	return node.pre != nil
}

// IsNil 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

// AddNodeFromHead 添加节点到链表头部的第N个元素之前, N=0表示新结点成为新的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n > list.len {
		panic("index out")
	}

	node := list.head

	for i := 1; i <= n; i++ {
		node = node.next
	}

	newNode := new(ListNode)
	newNode.value = v

	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		pre := node.pre
		if pre.IsNil() {
			newNode.next = node
			node.pre = newNode
			list.head = newNode
		} else {
			pre.next = newNode
			newNode.pre = pre
			node.next.pre = newNode
			newNode.next = node.next
		}
	}
	list.len += 1
}

// AddNodeFromTail 添加节点到链表尾部的第N个元素之后，N=0表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n > list.len {
		panic("index out")
	}

	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	newNode := new(ListNode)
	newNode.value = v

	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		next := node.next

		if next.IsNil() {
			node.next = newNode
			newNode.pre = node
			list.tail = newNode
		} else {
			newNode.pre = node
			node.next = newNode
			newNode.next = next
			next.pre = newNode
		}
	}
	list.len += 1
}

// IndexFromHead 从头部开始往后找，获取第N+1个位置的节点，索引0开始
func (list *DoubleList) IndexFromHead(n int) *ListNode {
	if n >= list.len {
		return nil
	}
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}
	return node
}

// IndexFromTail 从尾部开始往前找，获取第N+1个位置的节点，索引从0开始
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	if n >= list.len {
		return nil
	}
	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}
	return node
}

// PopFromHead 从头部开始往后找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n >= list.len {
		return nil
	}

	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	pre := node.pre
	next := node.next

	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		list.tail = pre
		pre.next = nil
	} else {
		pre.next = next
		next.pre = pre
	}
	list.len += 1
	return node
}

// PopFromTail 从尾部开始往前找，获取第N+1个位置的节点，并移除返回
func (list *DoubleList) PopFromTail(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()
	if n >= list.len {
		return nil
	}

	node := list.tail
	for i := 1; i <= n; i++ {
		node = node.pre
	}

	pre := node.pre
	next := node.next

	if pre.IsNil() && next.IsNil() {
		list.head = nil
		list.tail = nil
	} else if pre.IsNil() {
		list.head = next
		next.pre = nil
	} else if next.IsNil() {
		list.tail = pre
		pre.next = nil
	} else {
		pre.next = next
		next.pre = pre
	}
	list.len -= 1
	return node
}

func main() {

}
