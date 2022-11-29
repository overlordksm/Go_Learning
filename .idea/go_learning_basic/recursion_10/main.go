package main

import "fmt"

func recursion_test(n int) {
	if n > 2 && n < 10 {
		fmt.Println(n)
	} else {
		fmt.Println(n * 2)
	}
}
func main() {
	for i := 1; i < 15; i++ {
		recursion_test(i)
	}
}
