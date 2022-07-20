package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "John Doe"
	var age = 34
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

	fmt.Printf("%s is %d years old\n", name, age)
}
