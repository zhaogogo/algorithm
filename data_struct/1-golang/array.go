package main

import "fmt"

func main() {
	// TODO: 增加
	// 数组末尾追加（append）元素
	static_array_append()
	fmt.Println("==============================")
	// 数组中间插入（insert）元素
	static_array_insert()
	fmt.Println("==============================")
	// 数组空间已满
	static_array_expansion()
	fmt.Println("==============================")
	// TODO: 删除
	// 删除末尾元素
	static_array_pop()
	fmt.Println("==============================")
	// 删除中间元素
	static_array_delete()
	fmt.Println("==============================")
	// 动态数组, slice
	dynamic_array()
}

// 初始化数组，在数组末尾添加一个元素
func static_array_append() {
	// 创建大小为10的数组
	var arr [10]int
	// 数组中插入4个元素
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	// 在数组末尾添加一个元素
	arr[4] = 100
	// 打印数组
	fmt.Println(arr) // [0 1 2 3 100 0 0 0 0 0]
}

// 数组中间插入（insert）元素
func static_array_insert() {
	/*
	   比方说，我有一个大小为 10 的数组 arr，前 4 个索引装了元素，现在想在第 3 个位置（索引 2 arr[2]）插入一个新元素，怎么办？
	   这就要涉及「数据搬移」，给新元素腾出空位，然后再才能插入新元素。
	*/

	// 大小为 10 的数组已经装了 4 个元素
	var arr [10]int
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	fmt.Println("static_array_insert: ", arr) // [0 1 2 3 0 0 0 0 0 0]

	// 在索引 2 置插入元素 666
	// 需要把索引 2 以及之后的元素都往后移动一位
	// 注意要倒着遍历数组中已有元素避免覆盖，
	for i := 4; i < 2; i-- {
		arr[i] = arr[i-1]
	}
	// 现在第 3 个位置空出来了，可以插入新元素
	arr[2] = 666

	fmt.Println("static_array_insert: ", arr) // [0 1 666 3 0 0 0 0 0 0]
}

// 数组空间已满, 扩容静态数组
func static_array_expansion() {
	// 大小为 10 的数组已经装满了
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println("static_array_expansion: ", arr)
	// 现在想在数组末尾追加一个元素 10, 需要先扩容数组
	var arrNew [20]int
	for i := 0; i < len(arr); i++ {
		arrNew[i] = arr[i]
	}
	// 释放旧数组的内存空间
	// ...
	arrNew[10] = 10
	fmt.Println("static_array_expansion: ", arrNew)
}

// 删除末尾元素
func static_array_pop() {
	// 大小为 10 的数组已经装了 5 个元素
	var arr [10]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	fmt.Println("static_array_pop: ", arr)
	// 删除末尾元素，暂时用int默认类型0 代表元素已删除
	arr[4] = 0
	fmt.Println("static_array_pop: ", arr)
}

// 删除中间元素
func static_array_delete() {
	// 大小为 10 的数组已经装了 5 个元素
	var arr [10]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	fmt.Println("static_array_delete: ", arr)
	// 删除 arr[1]
	// 需要把 arr[1] 之后的元素都往前移动一位
	// 注意要正着遍历数组中已有元素避免覆盖
	for i := 1; i < 5; i++ {
		arr[i] = arr[i+1]
	}
	// // 最后一个元素置为0 代表已删除
	arr[4] = 0
	fmt.Println("static_array_delete: ", arr)
}

func dynamic_array() {
	// 创建动态数组
	// 不用显式指定数组大小，它会根据实际存储的元素数量自动扩缩容
	var slice []int = make([]int, 0)
	for i := 0; i < 10; i++ {
		// 在末尾追加元素，时间复杂度 O(1)
		slice = append(slice, i)
	}
	fmt.Println("\tdynamic_array: ", slice)
	// 在中间插入元素，时间复杂度 O(N)
	// 在索引 2 的位置插入元素 666
	slice = append(slice[:2], append([]int{666}, slice[3:]...)...)
	fmt.Println("在索引2的位置插入元素 666")
	fmt.Println("\tdynamic_array: ", slice)

	// 在头部插入元素，时间复杂度 O(N)
	fmt.Println("在头部插入元素，时间复杂度 O(N)")
	slice = append([]int{666}, slice...)
	fmt.Println("\tdynamic_array: ", slice)

	// 删除末尾元素，时间复杂度 O(1)
	fmt.Println("删除末尾元素，时间复杂度 O(1)")
	slice = slice[:len(slice)-1]
	fmt.Println("\tdynamic_array: ", slice)

	// 删除中间元素（索引 2 的元素），时间复杂度 O(N)
	fmt.Println("删除中间元素（索引 2 的元素），时间复杂度 O(N)")
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("\tdynamic_array: ", slice)

	// 根据索引[2]查询元素，时间复杂度 O(1)
	a := slice[2]
	fmt.Println("dynamic_array[2]: ", a)
	// 根据索引[2]修改元素，时间复杂度 O(1)
	slice[2] = 100
	fmt.Println("dynamic_array[2]: ", slice)
	// 根据元素值=666查找索引，时间复杂度 O(N)
	index := -1
	for i, v := range slice {
		if v == 666 {
			index = i
			break
		}
	}
	fmt.Println("dynamic_array 值为666的索引为: ", index)
}
