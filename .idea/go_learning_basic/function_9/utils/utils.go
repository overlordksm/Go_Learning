package utils

import "fmt"

var N1 float64
var N2 float64

func init() {
	N1 = 1.1
	N2 = 2.2
	fmt.Println("Keep this in mind, I(init) will be executed before main, you could use me for initialization works")
}

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
