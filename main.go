package main

import (
	"fmt"
	// "strconv"
)

func main() {
	//user input
	var identity string
	fmt.Print("Enter ur name ")
	fmt.Println(&identity) //It requires address of that variable -- pointers concepts
	fmt.Println(identity)
	fmt.Println("Welcome " + identity)

	//variables and how to print a combination of sentence using Print and Println

	// var name string = "Priya Pal"
	// age := 25
	// designation := "Software Engineer"
	// description := "learning go"
	// fmt.Println("Hi i am," + name +
	// 	" my age is " + strconv.Itoa(age) +
	// 	" i am a " + designation +
	// 	" Currently i am " + description)
	// fmt.Printf("Hi, my name is %s and my age is %d . I am a %s and at present %s\n",
	// 	name,
	// 	age,
	// 	designation,
	// 	description)

}
