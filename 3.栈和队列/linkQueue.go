package main

import "sync"

type LinkQueue struct {
	root *LinkNode2
	size int
	lock sync.Mutex
}

type LinkNode2 struct {
	Next  *LinkNode2
	Value string
}

// 入队

func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode2)
		queue.root.Value = v
	} else {
		newNode := new(LinkNode2)
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

func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	topNode := queue.root
	v := topNode.Value

	queue.root = topNode.Next
	queue.size -= 1

	return v
}

func main() {

}
