package main

import "fmt"

func oddEven(number int) {
	if number%2 == 0 {
		fmt.Println("Number is even :", number)
	} else {
		fmt.Println("Number is odd :", number)

	}
}
