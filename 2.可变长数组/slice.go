package main

import (
	"fmt"
	"sync"
)

type Array struct {
	array []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len   int        // 真正长度
	cap   int        // 容量
	look  sync.Mutex // 为了并发安全使用的锁
}

// 初始化数组

func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	// 把切片当数组用
	array := make([]int, cap, cap)

	//	元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// 增加一个元素

func (a *Array) Append(element int) {
	//	并发锁
	a.look.Lock()
	defer a.look.Unlock()

	//	大小等于容量，表示没有多余位置了
	if a.len == a.cap {
		//	没容量，数组要扩容，扩容到两倍
		newCap := 2 * a.len
		//	如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]int, newCap, newCap)

		//	把老数组的数据移动到新数组
		for k, v := range a.array {
			newArray[k] = v
		}

		//	替换数组
		a.array = newArray
		a.cap = newCap
	}

	//	把元素放在数组里
	a.array[a.len] = element
	// 真实长度+！
	a.len = a.len + 1
}

// 添加多个元素

func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// 获取指定下标元素

func (a *Array) Get(index int) int {
	//	如果越界了
	if a.len == 0 || index >= a.len {
		panic("index over len")
	}
	return a.array[index]
}

// 返回真实长度

func (a *Array) Len() int {
	return a.len
}

// 返回真实容量

func (a *Array) Cap() int {
	return a.cap
}

// 打印数组

func Print(array *Array) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		//	第一个元素
		if i == 0 {
			result = fmt.Sprintln("%s%d", result, array.Get(i))
			continue
		}

		result = fmt.Sprintln("%s%d", result, array.Get(i))
	}
	result = result + "]"
	return
}

func main() {
	//	创建一个容量为3的动态数组
	a := Make(0, 3)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	//	增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
	//	增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
