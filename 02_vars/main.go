package main

import "fmt"

func main() {
	size := 6.2

	name, email := "Brandon", "brandon_n_evans@live.com"

	var age int32 = 24
	var isCool = true
	isCool = false

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
