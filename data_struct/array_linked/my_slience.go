package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

const INITCAP = 1

var DEBUG = "1"

func init() {
	DEBUG = os.Getenv("DEBUG")
}

type MyType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

type MySlience[T MyType] struct {
	// 真正存储数据的底层数组
	data []T
	// 记录当前元素的个数
	size int
}

func NewMySlience[T MyType]() *MySlience[T] {
	return NewMySlienceWithCapacity[T](INITCAP)
}

func NewMySlienceWithCapacity[T MyType](size int) *MySlience[T] {
	return &MySlience[T]{
		data: make([]T, size),
		size: 0,
	}
}

// Append 增
// 数组末尾添加元素
func (l *MySlience[T]) Append(value T) {
	rowCap := len(l.data)
	// 判断底层数组容量沟通不够
	if rowCap == l.size {
		l.resize(2 * rowCap)
	}
	// 在尾部插入元素
	l.data[l.size] = value
	l.size++
}

func (l *MySlience[T]) Insert(index int, value T) error {
	// 检查索引越界
	if err := l.checkPositionIndex(index); err != nil {
		return err
	}
	// 判断是否需要扩容
	rowCap := len(l.data)
	if rowCap == l.size {
		l.resize(2 * rowCap)
	}
	// 迁移数据 data[index..] -> data[index+1..]
	for i := l.size - 1; i >= index; i-- { // TODO: *****
		l.data[i+1] = l.data[i]
	}
	// 插入元素
	l.data[index] = value

	l.size++
	return nil
}
func (l *MySlience[T]) AddFirst(value T) error {
	return l.Insert(0, value)
}

// Remove 删
func (l *MySlience[T]) Remove(index int) (T, error) {
	var empT T
	// 检查索引位置是否存在元素
	if l.size == 0 {
		return T(0), errors.New("MySlience not element")
	}
	if err := l.checkElementIndex(index); err != nil {
		return T(0), err
	}
	rowCap := len(l.data)
	// 缩容，节约空间
	if rowCap == l.size/4 {
		l.resize(rowCap / 2)
	}
	deleteVal := l.data[index]
	// 删除最后一个元素
	if l.size == index { // TODO: ****
		l.data[l.size-1] = empT
		l.size--
		return deleteVal, nil
	}

	// 迁移数据
	for i := index + 1; i < l.size; i++ { // TODO: ***
		l.data[i-1] = l.data[i]
	}
	l.data[l.size-1] = empT
	l.size--
	return deleteVal, nil
}
func (l *MySlience[T]) RemoveFirst() (T, error) {
	return l.Remove(0)
}
func (l *MySlience[T]) RemoveLast() (T, error) {
	return l.Remove(l.size - 1)
}

// Get 查
func (l *MySlience[T]) Get(index int) (T, error) {
	if err := l.checkElementIndex(index); err != nil {
		return T(0), err
	}
	return l.data[index], nil
}

// 改

func (l *MySlience[T]) Set(index int, value T) (T, error) {
	if err := l.checkElementIndex(index); err != nil {
		return T(0), err
	}
	oldValue := l.data[index]
	l.data[index] = value
	return oldValue, nil
}

// Size 工具方法============
func (l *MySlience[T]) Size() int {
	return l.size
}

func (l *MySlience[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *MySlience[T]) String() string {
	str := bytes.NewBuffer(nil)
	str.WriteString("MySlience size: ")
	str.WriteString(fmt.Sprintf("%v", l.size))
	str.WriteString(" , value: [ ")
	for i := 0; i < l.size; i++ {
		str.WriteString(fmt.Sprintf("%v ", l.data[i]))
	}
	str.WriteString("]")
	str.WriteString(" , raw_value: ")
	str.WriteString(fmt.Sprint(l.data))
	return str.String()
}

// ==================
// 将 data 的容量改为 newCap
func (l *MySlience[T]) resize(newSize int) {
	if DEBUG != "" {
		fmt.Printf("MySlience resize %d -> %d\n", l.size, newSize)
	}
	tmp := make([]T, newSize)
	for i := 0; i < l.size; i++ {
		tmp[i] = l.data[i]
	}
	l.data = tmp
}

// 检查 index 索引位置是否可以添加元素
func (l *MySlience[T]) checkPositionIndex(index int) error {
	if !l.isPositionIndex(index) {
		return errors.New(fmt.Sprintf("不允许插入的索引比底层数组索引还大, Input Index: %d, MySlience Size: %d", index, l.size))
	}
	return nil
}
func (l *MySlience[T]) isPositionIndex(index int) bool {
	// 为什么是 <= 判断 index <= l.size，
	// 索引是左闭右开的(], 等于size表示在数组的末尾添加元素
	return index >= 0 && index <= l.size
}

// 检查index索引位置是否存在元素
func (l *MySlience[T]) checkElementIndex(index int) error {
	if !l.isElementIndex(index) {
		return errors.New(fmt.Sprintf("索引位置不存在元素, Input Index: %d, MySlience size: %d", index, l.size))
	}
	return nil
}
func (l *MySlience[T]) isElementIndex(index int) bool {
	return index >= 0 && index < l.size
}

func main() {
	arr := NewMySlienceWithCapacity[int](3)
	for i := 0; i < 6; i++ {
		arr.Append(i)
	}
	fmt.Println(arr)
	v, err := arr.Remove(5) // 删除的值: 5 <nil>
	//v, err := arr.Remove(6)  // 删除的值: 0 索引位置不存在元素, Input Index: 6, MySlience size: 6
	fmt.Println("删除的值:", v, err)
	fmt.Println(arr)
	err = arr.AddFirst(100)
	fmt.Println(arr, err)
	err = arr.AddFirst(101)
	fmt.Println(arr, err)
	_, err = arr.RemoveLast()
	fmt.Println(arr, err)
}
