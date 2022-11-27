package main

import "fmt"

func main() {
	// float result type
	fmt.Println(10 / 4)

	fmt.Println(10.0 / 4)

	var n float32 = 10 / 4

	fmt.Println(n)

	// %
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(-10 % -3)

	// ++, --
	aa := 11
	aa++

	fmt.Println(aa)

	aa--

	fmt.Println(aa)

	// relational operator
	fmt.Println(1 == 1)
	fmt.Println(1 > 1)
	fmt.Println(1 < 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 != 1)

	// logical operator
	fmt.Println(1 == 1 && 1 != 1)
	fmt.Println(1 == 1 || 1 != 1)
	if 1 == 1 || 1 != 1 {
		fmt.Println("1 == 1 or 1 != 1")
	}

	//
}
