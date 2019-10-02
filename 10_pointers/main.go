package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read val from address
	*b = 10
	fmt.Println(*b)
}
