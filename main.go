package main

import (
	"fmt"
	"gopractice/calc"
	"gopractice/loops"
	"gopractice/pointers"
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

	loops.TableMultiplication(7)
	//implementing user input to give specific output based on user preference helps.
	var num int
	fmt.Print("Enter a number to generate its multiplication table: ")
	_, err := fmt.Scan(&num) //err refers to error in Go and error is stored in variable err, '_' (underscore) is a blank identifier that discards the unused value. By using _, you tell Go that you’re only interested in the error, not the number of scanned items.
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	loops.TableMultiplication(num)

	loops.SumOfElements([]int{0, 1, 2, 3, 4, 5, 6})

	sum = loops.SumOfElements([]int{0, 1, 2, 3, 4, 5, 6})
	fmt.Println("Total sum of arr3:", sum)

	loops.PrimeNums()

	pointers.PointersSwapNums(14, 28)
	//pointers.DoubleTheValue(5)

}
