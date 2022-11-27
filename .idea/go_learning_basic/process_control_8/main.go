package main

import "fmt"

func main() {
	// if demo
	var age int

	fmt.Println("Please enter age: ")
	fmt.Scanln(&age)
	fmt.Printf("age is: %v \n", age)
	if age >= 18 {
		fmt.Println("You should be accountable for your behaviors")
	}

	if age := 20; age >= 18 {
		fmt.Println("Your age is older than 18, you should be accoutable for your behaviors")
	}

	// if-else demo
	var salary float32

	fmt.Println("Please enter your salary: ")
	fmt.Scanln(&salary)

	if salary > 30000 {
		fmt.Println("You should pay some taxes.")
	} else {
		fmt.Println("You don't have to pay any tax")
	}

	// if-else if-else demo
	var ads_score float32

	fmt.Println("Please enter the Ads Score of your ranking model: ")
	fmt.Scanln(&ads_score)

	if ads_score >= 80 {
		fmt.Println("Your model is doing pretty well!")
	} else if ads_score >= 60 {
		fmt.Println("Your model is working fine")
	} else {
		fmt.Println("Hey dude, what's wrong with it? Better create a SEV to trouble shoot it!")
	}

	// switch demo
	var favorite_language string

	fmt.Println("Plese enter your favorite language(capital letter please): ")
	fmt.Scanln(&favorite_language)

	switch favorite_language {
	case "C":
		fmt.Println("System Development guy?")
	case "C++":
		fmt.Println("Highly performance focused guy?")
	case "python":
		fmt.Println("AI guy?")
	case "SQL":
		fmt.Println("Data Guy?")
	case "Golang":
		fmt.Println("Smart guy!")
	default:
		fmt.Println("???")
	}

	// for loop demo
	for i := 1; i <= 10; i++ {
		if i > 3 {
			fmt.Println(i)
		}
	}

	//
}
