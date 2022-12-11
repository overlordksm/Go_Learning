package main

// alias for the library name
import (
	u "Go_Learning/.idea/go_learning_basic/function_9/utils"
	"fmt"
)

func init() {
	fmt.Println("Keep this in mind, I(init) will be executed before main, you could use me for initialization works")
}

func main() {
	u.Calculate(u.N1, u.N2, '+')

	u.Calculate(u.N1, u.N2, '-')

	u.Calculate(u.N1, u.N2, '*')

	u.Calculate(u.N1, u.N2, '/')

	a := func(n1 float64, n2 float64) float64 {
		return n1 * n2
	}

	fmt.Println(a(u.N1, u.N2))
}
