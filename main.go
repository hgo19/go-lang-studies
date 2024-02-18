package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"guts", "joe", "thorfinn", "saitama", "mashle"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at position", index)
			continue // go back to the for loop and do not execute anything below this
		}

		if index > 2 {
			fmt.Println("breaking at position", index)
			break // stop the for loop
		}

		fmt.Printf("the value at position %v is %v \n", index, value)
	}

}