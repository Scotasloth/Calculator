package main

import (
	"fmt" // For input and output
	"log" // For logging errors
	// For converting string inputs to numbers
)

func main() {

	var num1, num2 float64
	var operation string

	fmt.Print("Welcome")

	fmt.Print("Please enter the first number: ")

	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter the second number: ")

	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Select operator: + - * /: ")
	fmt.Scan(&operation)

	var result float64
	switch operation {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			log.Fatal("Error division by 0")
		}
	default:
		log.Fatal("Invalid Operation")
	}

	fmt.Printf("Result: %v\n", result)
}
