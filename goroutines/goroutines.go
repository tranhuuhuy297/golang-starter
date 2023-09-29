package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 100)
	}
}

func sayWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go sayHello()
	go sayWorld()
	time.Sleep(time.Second)
}
