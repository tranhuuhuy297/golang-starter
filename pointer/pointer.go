package main

import "fmt"

func main() {
	var p *int
	i := 4
	p = &i

	fmt.Println(p, i)
	*p = 21
	fmt.Println(i)
}
