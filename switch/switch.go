package main

import "fmt"

func main() {
	number := 5
	switch number {
	case 1:
		fmt.Print(1)
	case 3, 5:
		fmt.Print(3)
	default:
		fmt.Print("default switch")
	}
}
