package main

import "fmt"

func main() {
	a := '3'
	b := 3
	// if a == b {
	// 	fmt.Println("a == b")
	// }
	if int(a) == b {
		fmt.Println("a == b")
	} else if b == 3 {
		fmt.Println("test else if")
	} else {
		fmt.Println(a, b)
	}
}
