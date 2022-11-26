package main

import "fmt"

func main() {
	var i int = 10

	fmt.Println("the address of i is: ", &i)

	// 1. ptr is a pointer variable
	// 2. the tyupe of ptr is *int
	// 3. the value of ptr is &i
	var ptr *int = &i
	fmt.Println(ptr)

	var n int = 9
	fmt.Println("address of n is:", &n)

	var ptr2 *int = &n
	fmt.Println("address of n is:", ptr2)
}
