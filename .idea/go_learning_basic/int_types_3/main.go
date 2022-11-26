package main

import (
	"fmt"
	"unsafe"
)

func main(){
	// int8, 8 means 2^8, range from -2^7 to 2^7 - 1
	var a int = 1
	var b uint = 2
	var c byte = 255
	d := 5.1234E2
	e := .12
	var char1 byte = 'a'
	var char2 byte = '0'
	var bl bool = false
	var Default int

	fmt.Println("a = ", a, "\nb = ", b, "\nc = ", c, "\nsize of a / b / c are ", unsafe.Sizeof(a), " ", unsafe.Sizeof(b), " ", unsafe.Sizeof(c))
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

	// will print ACSII codes when char are saved as bytes
	fmt.Println("char1 = ", char1)
	fmt.Println("char2 = ", char2)

	// print with char formats
	fmt.Printf("char1 = %c", char1)
	fmt.Printf("\nchar2 = %c", char2)

	// bool example
	fmt.Println("\nbl = ", bl)

	// default values
	fmt.Println("Default = ", Default)

	// data type transformation

	var dtt int = 10

	fmt.Println(float32(dtt) + .01)
}