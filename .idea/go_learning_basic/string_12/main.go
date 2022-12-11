package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// len of strings will return the length in bytes not in characters
	str := "hello北"
	fmt.Println("str len = ", len(str))

	r := []rune(str)
	for i := 0; i < len(str)-2; i++ {
		fmt.Printf("%c\n", r[i])
	}

	n, err := strconv.Atoi("123") // strconv.Atoi("hello")
	if err != nil {
		fmt.Println("transformation error: ", err)
	} else {
		fmt.Println(n)
	}

	i := strconv.Itoa(12345)
	fmt.Printf("str = %v, str = %T\n", i, i)

	s1 := strconv.FormatInt(12345, 2)
	s2 := strconv.FormatInt(12345, 16)
	fmt.Println(s1, s2)

	fmt.Println(strings.Count("cheese", "e"))
}
