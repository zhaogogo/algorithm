package main

import "fmt"

func main() {
	var a = "我哎"
	for _, v := range a {
		fmt.Printf("%T  %v\n", v, v)
	}
}
