package utils

import "fmt"

func Calculate(n1 float64, n2 float64, operator byte) {
	var res float64

	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("Invalid Input")
	}

	fmt.Println(res)
}
