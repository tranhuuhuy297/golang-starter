package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go print(&wg)

	wg.Wait()

	fmt.Println("done")
}

func print(wg *sync.WaitGroup) { // Phải truyền pointer vào để tránh việc truyền bản sao ->
	fmt.Println("hello from goroutine")
	wg.Done()
}
