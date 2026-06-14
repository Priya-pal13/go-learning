package main

import "fmt"

func addition(a int, b int) int {
	return a + b

}
func subtraction(a int, b int) int {
	return a - b

}
func multiplication(a int, b int) int {
	return a * b

}
func division(a int, b int) float64 {
	if b == 0 {
		fmt.Println("Cannot divided by 0")
		return 0
	}
	return float64(a) / float64(b)

}

var num1 int
var num2 int
var operator string

func calculator(a int, b int) {

	fmt.Println("Enter your operator +,-,/,*  :")
	fmt.Scanln(&operator)

	switch operator {

	case "+":

		result := addition(a, b)
		fmt.Println("Result :", result)

	case "-":
		result := subtraction(a, b)
		fmt.Println("Result :", result)

	case "*":

		result := multiplication(a, b)
		fmt.Println("Result :", result)

	case "/":

		result := division(a, b)
		fmt.Println("Result :", result)

	default:

		fmt.Println("Invalid operator")

	}
}
