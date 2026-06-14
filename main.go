package main

func main() {
	//user input
	// var identity string
	// fmt.Print("Enter ur name ")
	// fmt.Println(&identity) //It requires address of that variable -- pointers concepts
	// fmt.Println(identity)
	// fmt.Println("Welcome " + identity)

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

	//take user input and calculating the sum of those numbers
	// var num int
	// sum := 0
	// var i int
	// for i = 0; i < 5; i++ {
	// 	fmt.Println("Please enter your number")
	// 	fmt.Scanln(&num)
	// 	sum += num
	// }
	// fmt.Println("Total sum is " + strconv.Itoa(sum))

	// take an user input to check whether you are eligible to vote
	// var name string
	// var age int
	// fmt.Println("Enter your name : ")
	// fmt.Scanln(&name)
	// fmt.Println("Enter your age : ")
	// fmt.Scanln(&age)
	// if age >= 18 {
	// 	fmt.Println("You are eligible to vote")
	// } else {
	// 	fmt.Println("You are not eligible to vote")
	// }

	//Even and odd checks
	// var number int
	// fmt.Println("Enter a number : ")
	// fmt.Scanln(&number)

	// if number%2 == 0 {
	// 	fmt.Println("Your number : " + strconv.Itoa(number) + " is a even number")
	// } else {
	// 	fmt.Println("Your number : " + strconv.Itoa(number) + " is a odd number")

	// }

	// ask user for 5 number and check how many are odd and how many are even
	// var number int

	// var evennum []int
	// var oddnum []int

	// var i int
	// for i = 0; i < 5; i++ {

	// 	fmt.Println("Enter your numbers : ")
	// 	fmt.Scanln(&number)
	// 	if number%2 == 0 {
	// 		fmt.Println("Your number : " + strconv.Itoa(number) + " is a even number")
	// 		evennum = append(evennum, number)

	// 	} else {
	// 		fmt.Println("Your number : " + strconv.Itoa(number) + " is a odd number")
	// 		oddnum = append(oddnum, number)

	// 	}
	// 	fmt.Println("The even numbers are : ", evennum)
	// 	fmt.Println("The odd numbers are : ", oddnum)

	// }
	// fmt.Println("Count of the even numbers", len(evennum))
	// fmt.Println("Count of the odd numbers", len(oddnum))

	// Numbers: [10 5 8 3 12]
	// Sum: 38
	// Average: 7.6
	// Max: 12
	// Min: 3
	// Even count: 3
	// Odd count: 2
	//need to show all the above

	// var i int
	// var number int
	// var Numbers []int
	// var sum int
	// var average int
	// var count int = 0

	// var EvenCount int = 0
	// var OddCount int = 0
	// var max int
	// var min int
	// for i = 0; i < 5; i++ {
	// 	fmt.Println("Enter the numbers")
	// 	fmt.Scanln(&number)
	// 	Numbers = append(Numbers, number)
	// 	count++
	// 	sum += number
	// 	average = sum / count
	// 	if number%2 == 0 {
	// 		EvenCount++
	// 	} else {
	// 		OddCount++
	// 	}

	// 	//min - max

	// 	if i == 0 {
	// 		max = number
	// 		min = number
	// 	} else {
	// 		if number > max {
	// 			max = number
	// 		}
	// 		if number < min {
	// 			min = number
	// 		}
	// 	}
	// 	fmt.Println("Numbers are : ", Numbers)
	// 	fmt.Println("Sum is :", sum)
	// 	fmt.Println("Average is ", average)
	// 	fmt.Println("Even count is ", EvenCount)
	// 	fmt.Println("Odd count is ", OddCount)
	// 	fmt.Println("Max number is", max)
	// 	fmt.Println("Min number is ", min)
	// }

	//Second Largest number

	// LargestNumber := 0
	// SecondLargestNumber := 0
	// var i int
	// var number int
	// for i = 0; i < 5; i++ {
	// 	fmt.Println("Enter the numbers")
	// 	fmt.Scanln(&number)
	// 	if i == 0 {
	// 		LargestNumber = number
	// 		SecondLargestNumber = number
	// 	}
	// 	if number > LargestNumber {
	// 		SecondLargestNumber = LargestNumber
	// 		LargestNumber = number
	// 	} else if number > SecondLargestNumber && number != LargestNumber {
	// 		SecondLargestNumber = number
	// 	}
	// }
	// fmt.Println("Largest number is ", LargestNumber)

	// fmt.Println("Second Largest number is ", SecondLargestNumber)

	//calling the calculate_sum.go
	add(12, 23)
	//odd-even function
	oddEven(167)
	calculator(2, 7)
}
