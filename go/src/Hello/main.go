package main

import "fmt"

func main() {

	challenge := "#90DaysOfDevOps"
	const daysTotal int = 90
	var remainingDays uint = 90

	fmt.Printf("Welcome to %v\n challenge.\nThis challenge consists of %v days\n", challenge, daysTotal)

	var twitterName string
	var daysCompleted uint
	// asking for user input

	fmt.Println("Enter Your Twitter Handle: ")
	fmt.Scanln(&twitterName)

	fmt.Println("How many days have you completed?: ")
	fmt.Scanln(&daysCompleted)

	// calculate remaining days
	remainingDays = remainingDays - daysCompleted

	fmt.Printf("Thank you %v for taking part and completing %v days.\n", twitterName, daysCompleted)
	fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
	fmt.Println("Good luck")
}
