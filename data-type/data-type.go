package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 3
	fmt.Println(reflect.TypeOf(a))

	var var_string string = "string"
	var var_int int = 3
	var var_float float32 = 3.5
	var var_bool bool // 1 or true = true

	fmt.Scan(&var_bool)

	fmt.Println("%v, %v, %v, %v", var_string, var_int, var_float, var_bool)
}
