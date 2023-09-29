package main

import "fmt"

type User struct {
	name       string
	age        int
	isActivate bool
}

func main() {
	var user1 = User{
		name:       "huy",
		age:        15,
		isActivate: false,
	}
	fmt.Println(user1)
}
