package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AskForNumber() int64 {
	// from the bufio module we make
	// a newScanner object that has
	// os standart data input  Stdin
	scanner := bufio.NewScanner(os.Stdin)
	// Prompt message
	fmt.Printf("Give me a number: ")
	// Scans the line and stores
	// the data in the scanner object
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || input < 0 {
		fmt.Println("Incorrect input!")
		os.Exit(1)
	}
	return input
}

func AskForOperation() int64 {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your desired operation: \n1) addition\n2) subtraction\n3) division\n4) multiplication\n")
	fmt.Printf("Choose the numeric representation of the operation: ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || input > 4 {
		fmt.Println("Incorrect input!")
		os.Exit(1)
	}
	return input
}

func CheckBiggerNumber(x, y int64) (z1, z2 int64) {
	if x > y {
		z1 = x
		z2 = y
	} else {
		z1 = y
		z2 = x
	}
	return
}

func main() {
	firstNumber := AskForNumber()
	secondNumber := AskForNumber()
	operation := AskForOperation()
	num1, num2 := CheckBiggerNumber(firstNumber, secondNumber)
	var result int64

	switch operation {
	case 1:
		result = (num1 + num2)
		fmt.Printf("The result is: %d\n", result)
	case 2:
		result = (num1 - num2)
		fmt.Printf("The result is: %d\n", result)
	case 3:
		result = (num1 / num2)
		fmt.Printf("The result is: %d\n", result)
	case 4:
		result = (num1 * num2)
		fmt.Printf("The result is: %d\n", result)
	}
}
