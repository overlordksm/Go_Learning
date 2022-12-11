package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("%T %v, %v \n", num1, num1, &num1)

	num2 := new(int64)
	*num2 = 100
	fmt.Printf("%T %v, %v, %v \n", num2, num2, &num2, *num2)
	// num2's type is *int
	// num2's value is the address of num2 pointer variable
	// the pointer address points to num2 is &num2
	// num2's points to value *num2
}
