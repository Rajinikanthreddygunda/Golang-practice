package main

import (
	"fmt"
	"gopractice/calc"
	"gopractice/loops"
)

func main() {
	sum := calc.Add(2, 3)
	fmt.Println("sum of a+b:", sum)

	sub := calc.Subtract(20, 3)
	fmt.Println("subtraction of a, b:", sub)

	Multiply := calc.Multiply(10, 5)
	fmt.Println("Multiply of a, b:", Multiply)

	arr1 := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// Call the Oddeven function and pass the array as an argument
	loops.Oddeven(arr1)

	arr2 := make([]int, 9, 20)

	// Initialize the slice with values
	arr2 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("slice function call")

	// Call the Oddeven function and pass the slice as an argument
	loops.SliceOddeven(arr2)
}
