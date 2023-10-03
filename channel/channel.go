package main

import (
	"fmt"
	"time"
)

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	fmt.Println(len(s))
// 	c <- sum
// }

// func main() {
// 	s := []int{1, 2, 3}
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c
// 	fmt.Println(x, y)
// }

// func test(s int, c chan int) {
// 	fmt.Println("send ", s)
// 	c <- s
// }

// func main() {
// 	c := make(chan int, 1)
// 	go test(3, c)
// 	go test(4, c)
// 	go test(5, c)
// 	go test(6, c)
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

func main() {
	myChan := make(chan int)
	go func() {
		fmt.Println(<-myChan)
	}()
	fmt.Println(4)
	myChan <- 1
	fmt.Println(3)
	time.Sleep(2) // sleep để main không end từ đoạn print 3
}
