package main

import "fmt"

// go lang không có khởi tạo giá trị mặc định cho param trong go lang

// go run function/*

func greet(name string) {
	fmt.Printf("%v\n", name)
	test()
}

func return_list(test []string) []string {
	return test
}

func main() {
	greet("huy")
	fmt.Println([]string{"trần", "hữu", "huy"})
}
