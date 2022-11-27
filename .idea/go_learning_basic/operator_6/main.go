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

	// =, +=, -=, *=, /=, %=
	var assign_val float64
	assign_val = 10.0

	assign_val += 1
	fmt.Println(assign_val)

	assign_val -= 1
	fmt.Println(assign_val)

	assign_val *= 10
	fmt.Println(assign_val)

	assign_val /= 10
	fmt.Println(test_func(assign_val))

	assign_val_int := int(assign_val)
	assign_val_int %= 3
	fmt.Println(assign_val_int)

	// position operators: &, |, ^, <<, >>
	// 2 & 3: 0000 0010 & 0000 0011 -> 0000 0010
	// 2 | 3: 0000 0010 | 0000 0011 -> 0000 0011
	// 2 ^ 3(异或，一个值为0， 一个值为1): 0000 0010 ^ 0000 0011 -> 0000 0001
	// -2 ^ 2: 1111 1110 ^ 0000 0010 -> 1111 1100 -> 1111 1011 -> 1000 0100
	// 2 << 3:
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 3)
	fmt.Println(2 << 3)
	fmt.Println(2.0 >> 3)
}

func test_func(val float64) float64 {
	return val * 10
}
