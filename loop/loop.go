package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
}
