package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 10)
	for _, v := range slice1 {
		fmt.Println(v)
	}
}
