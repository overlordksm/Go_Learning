package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	fmt.Println(year)

	month := now.Month()
	fmt.Println(month)

	day := now.Day()
	fmt.Println(day)

	hour := now.Hour()
	fmt.Println(hour)

	min := now.Minute()
	fmt.Println(min)

	fmt.Printf("%d-%d-%d %d: %d \n", year, month, day, hour, min)

	// unix system timestamp
	fmt.Println(now.Unix(), now.UnixNano())

	// time.Sleep()
	//for i:= 0; i < 100; i++ {
	//	fmt.Println(i)
	//	time.Sleep(time.Second / 10)
	//}
}
