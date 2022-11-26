package main
import (
	"fmt"
	"strconv"
	_ "strconv"
	_ "unsafe"
)
func main(){
	var n1 int = 99
	var n2 float64 = 23.456
	var n3 int = 99
	var n4 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	str = strconv.FormatInt(int64(n3), 10)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	str = strconv.FormatFloat(n4, 'f', 10, 64)
	fmt.Printf("str type %T str = %v", str, str)

	fmt.Println("\n")

	var str1 string = "true"
	var b1 bool

	b1, _ = strconv.ParseBool(str1)

	fmt.Println("b1 type is: %T b1=%v", b1)

	fmt.Println("\n")

	var str2 string = "123"
	var n5 int64

	n5, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n5 type %T n5=%v", n5, n5)

	var str3 string = "123.456"
	var n6 float64

	n6, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("n6 type %T n6=%v", n6, n6)
}
