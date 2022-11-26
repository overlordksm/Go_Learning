package main

import "fmt"

func main(){
	// Declare Variables
	var i int
	var j = 1.1
	k := 2

	var (
		a, b, c = 3, 4, 5.5
	)

	var str1 = "Hello "
	str2 := "Go!"

	// Assign value to i
	i = 10

	a = 11

	// use variable

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("a + b + c = ", float32(a) + float32(b) + float32(c))
	fmt.Println(str1 + str2)
}