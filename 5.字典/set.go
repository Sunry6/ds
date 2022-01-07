package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m   map[int]struct{} // 用字典来实现，因为字段键不能重复
	len int
	sync.Mutex
}

// NewSet 新建一个空集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{
		m: temp,
	}
}

// Add 增加一个元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

// Remove 删除一个元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return
	}

	delete(s.m, item)
	s.len = len(s.m)
}

// Has 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.m[item]
	return ok
}

// Len 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// IsEmpty 集合是否为空
func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// Clear 消除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

// List 将集合转化为列表
func (s *Set) List() []int {
	s.Lock()
	defer s.Unlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	s := NewSet(5)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	fmt.Println("list of all items", s.List())

	s.Clear()
	if s.IsEmpty() {
		fmt.Println("empty")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Has(2) {
		fmt.Println("2 does exist")
	}

	s.Remove(2)
	s.Remove(3)
	fmt.Println("list of all items", s.List())
}
