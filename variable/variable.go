package main

import "fmt"

func main() {
	var a = 3
	fmt.Println(a)

	var b, c, d = 4, 5, 6
	fmt.Printf("%d, %d, %d\n", b, c, d)

	const PI = 3.14
	// PI = 4 -> error
}
