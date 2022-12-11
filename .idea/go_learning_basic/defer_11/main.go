package main

import "fmt"

func sum(n1 int, n2 int) int {
	defer fmt.Println("test defer n1 = ", n1)
	defer fmt.Println("test defer n2 = ", n2)
	res := n1 + n2
	fmt.Println("test defer res = ", res)
	fmt.Println("The main reason of utilizing defer is to release function space timely")

	return res
}

func main() {
	sum(1, 2)
}
