package main

import "fmt"


var someName = "hello" // common uses offside of a function

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi" // infer the type
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"


	nameFour := "yoshi" // can't be used offside of a function

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)


	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40 // can't be used offside of a function

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25 // specific range of numbers 250 would be outside of the scope in this case (-128 through 127)
	var numTwo int8 = -128
	var numThree uint16 = 256 // 0 to 255
	fmt.Println(numOne, numTwo, numThree)


	// float

	var scoreOne float32 = 25.98 // can be positive or negative number
	var scoreTwo float64 = 48373738435453454343.7 // used in the most of the time

	scoreThree := 1.5
	
	fmt.Println(scoreOne, scoreTwo, scoreThree)
} 