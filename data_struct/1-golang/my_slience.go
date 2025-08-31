package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

const INIT_CAP = 1

var DEBUG = ""

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
	return NewMySlienceWithCapacity[T](INIT_CAP)
}

func NewMySlienceWithCapacity[T MyType](size int) *MySlience[T] {
	return &MySlience[T]{
		data: make([]T, size),
		size: 0,
	}
}

// 增
// 数组末尾添加元素
func (this *MySlience[T]) Append(value T) {
	cap := len(this.data)
	// 判断底层数组容量沟通不够
	if cap == this.size {
		this.resize(2 * cap)
	}
	// 在尾部插入元素
	this.data[this.size] = value
	this.size++
}

func (this *MySlience[T]) Insert(index int, value T) error {
	// 检查索引越界
	if err := this.checkPositionIndex(index); err != nil {
		return err
	}
	// 判断是否需要扩容
	cap := len(this.data)
	if cap == this.size {
		this.resize(2 * cap)
	}
	// 迁移数据 data[index..] -> data[index+1..]
	for i := this.size; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}
	// 插入元素
	this.data[index] = value

	this.size++
	return nil
}
func (this *MySlience[T]) AddFirst(value T) error {
	return this.Insert(0, value)
}

// 删
func (this *MySlience[T]) Remove(index int) (T, error) {
	var empT T
	// 检查索引位置是否存在元素
	if this.size == 0 {
		return empT, errors.New("MySlience not element")
	}
	if err := this.checkElementIndex(index); err != nil {
		return empT, err
	}
	cap := len(this.data)
	// 缩容，节约空间
	if cap == this.size/4 {
		this.resize(cap / 2)
	}
	deleteVal := this.data[index]
	// 删除最后一个元素
	if this.size == index {
		this.data[this.size-1] = empT
		this.size--
		return deleteVal, nil
	}

	// 迁移数据
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.data[this.size-1] = empT
	this.size--
	return deleteVal, nil
}
func (this *MySlience[T]) RemoveFirst() (T, error) {
	return this.Remove(0)
}
func (this *MySlience[T]) RemoveLast() (T, error) {
	return this.Remove(this.size - 1)
}

// 查
func (this *MySlience[T]) Get(index int) (T, error) {
	if err := this.checkElementIndex(index); err != nil {
		return T(0), err
	}
	return this.data[index], nil
}

// 改

func (this *MySlience[T]) Set(index int, value T) (T, error) {
	if err := this.checkElementIndex(index); err != nil {
		return T(0), err
	}
	oldValue := this.data[index]
	this.data[index] = value
	return oldValue, nil
}

// 工具方法============
func (this *MySlience[T]) Size() int {
	return this.size
}

func (this *MySlience[T]) IsEmpty() bool {
	return this.size == 0
}

func (this *MySlience[T]) String() string {
	str := bytes.NewBuffer(nil)
	str.WriteString("MySlience size: ")
	str.WriteString(fmt.Sprintf("%v", this.size))
	str.WriteString(" , value: [ ")
	for i := 0; i < this.size; i++ {
		str.WriteString(fmt.Sprintf("%v ", this.data[i]))
	}
	str.WriteString("]")
	str.WriteString(" , raw_value: ")
	str.WriteString(fmt.Sprint(this.data))
	return str.String()
}

// ==================
// 将 data 的容量改为 newCap
func (this *MySlience[T]) resize(newSize int) {
	if DEBUG != "" {
		fmt.Printf("MySlience resize %d -> %d\n", this.size, newSize)
	}
	tmp := make([]T, newSize)
	for i := 0; i < this.size; i++ {
		tmp[i] = this.data[i]
	}
	this.data = tmp
}

// 检查 index 索引位置是否可以添加元素
func (this *MySlience[T]) checkPositionIndex(index int) error {
	if !this.isPositionIndex(index) {
		return errors.New(fmt.Sprintf("不允许插入的索引比底层数组索引还大, Input Index: %d, MySlience Size: %d", index, this.size))
	}
	return nil
}
func (this *MySlience[T]) isPositionIndex(index int) bool {
	// 为什么是 <= 判断 index <= this.size，
	// 索引是左闭右开的(], 等于size表示在数组的末尾添加元素
	return index >= 0 && index <= this.size
}

// 检查index索引位置是否存在元素
func (this *MySlience[T]) checkElementIndex(index int) error {
	if !this.isElementIndex(index) {
		return errors.New(fmt.Sprintf("索引位置不存在元素, Input Index: %d, MySlience size: %d", index, this.size))
	}
	return nil
}
func (this *MySlience[T]) isElementIndex(index int) bool {
	return index >= 0 && index < this.size
}

func main() {
	arr := NewMySlienceWithCapacity[int](3)
	for i := 0; i < 6; i++ {
		arr.Append(i)
	}
	fmt.Println(arr)
	v, err := arr.Remove(5)
	fmt.Println("删除的值:", v, err)
	fmt.Println(arr)
}
