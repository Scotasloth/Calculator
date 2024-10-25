package main

import (
	"fmt"     // For input and output
	"log"     // For logging errors
	"strconv" // For converting string inputs to numbers
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

	strconv.AppendQuoteRuneToASCII() //here so it doesnt remove my import when saving

}
