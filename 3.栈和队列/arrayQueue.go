package main

import "sync"

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

// 入队

func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)
	queue.size += 1
}

// 出队

func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	v := queue.array[0]

	//for i := 1; i < queue.size; i++ {
	//	queue.array[i-1] = queue.array[i]
	//}
	//queue.array = queue.array[0 : queue.size-1]

	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	queue.size -= 1
	return v
}

func main() {

}
