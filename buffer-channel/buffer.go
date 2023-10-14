package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	fmt.Println("write func")
	close(ch)
}
func main() {

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("done wait")
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("done main")
}
