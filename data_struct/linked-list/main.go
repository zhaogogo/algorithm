package main

import (
	"fmt"
	"zhao/pkg"
)

func main() {
	link := pkg.NewSingleList()
	fmt.Println(link.Append(&pkg.SingleNode{Data: 0}))
	fmt.Println(link.Append(&pkg.SingleNode{Data: 1}))
	fmt.Println(link.Append(&pkg.SingleNode{Data: 2}))
	fmt.Println(link.Append(&pkg.SingleNode{Data: 3}))
	fmt.Println(link.Append(&pkg.SingleNode{Data: 4}))
	fmt.Println(link.Append(&pkg.SingleNode{Data: 5}))
	link.String()
	link.Delete(5)
	link.String()
}
