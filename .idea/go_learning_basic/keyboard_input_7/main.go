package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var salary float32
	var isPass bool

	fmt.Println("Please enter name: ")
	fmt.Scanln(&name)

	fmt.Println("Please enter age: ")
	fmt.Scanln(&age)

	fmt.Println("Please enter salary: ")
	fmt.Scanln(&salary)

	fmt.Println("Please enter if passing the test: ")
	fmt.Scanln(&isPass)

	fmt.Printf("name is %v \n age is %v \n salary is %v \n whether pass the test %v \n", name, age, salary, isPass)
}
