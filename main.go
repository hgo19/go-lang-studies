package main

import "fmt"

func main() {

	// Arrays
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // short hand

	names := [4]string{"saitama", "guts", "joe", "thorfinn"}
	names[1] = "griffith"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85) // append returns a new slice

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] // includes the first number but not the second
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "ashito")

	fmt.Println(rangeOne)
}