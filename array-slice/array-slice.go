package main

import "fmt"

func main() {
	b := []int{1, 2, 3}
	fmt.Println(b)
	b[0] = 4
	fmt.Println(b)
}
