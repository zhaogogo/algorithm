package pkg

import (
	"fmt"
	"sync"
)

type SingleObject interface{}

type SingleNode struct {
	Data SingleObject
	Next *SingleNode
}

type SingleList struct {
	mutex *sync.Mutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

func NewSingleList() *SingleList {
	return &SingleList{
		mutex: new(sync.Mutex),
		Head:  nil,
		Tail:  nil,
		Size:  0,
	}
}

func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}
	list.Tail.Next = node
	list.Tail = node
	list.Size += 1
	return true
}

func (list *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil {
		return false
	}
	if index > list.Size {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Size += 1
		return true
	}
	var i uint
	ptr := list.Head
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	oldNext := ptr.Next
	ptr.Next = node
	node.Next = oldNext
	list.Size += 1
	return true
}

func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Size == 0 || list.Size-1 < index {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		head := list.Head.Next
		list.Head = head
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}

	ptr := list.Head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next

	ptr.Next = next.Next
	if index == list.Size-1 {
		list.Tail = ptr
	}
	list.Size -= 1
	return true
}

func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.Size == 0 || list.Size-1 < index {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (list *SingleList) String() {
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	l := make([]interface{}, 0, list.Size)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++ {
		l = append(l, ptr.Data)
		ptr = ptr.Next
	}
	fmt.Println("linked_list: ", l)
}
