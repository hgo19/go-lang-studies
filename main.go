package main

import "fmt"

func main() {
	age := 35
	name := "saitama"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello mario world!")
	fmt.Println("goodbye mario world!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) %+ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55) // u can put the number of decimal places u want in the number

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}