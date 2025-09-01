package main

import (
	"container/list"
	"fmt"
)

func main() {
	list1 := list.New()
	a := list1.PushFront("a")
	fmt.Println(list1.Remove(a))
	fmt.Println(list1.Remove(a))

}
