package pointers

import "fmt"

//Task: Create a program to swap two numbers using pointers.

/*func PointersSwapNums(a, b int) {
	num1 := a
	fmt.Println("num1:", num1)
	num2 := &num1
	*num2 = b
	fmt.Println("updated num1:", num1)
} //ignore*/

func PointersSwapNums(a, b *int) {
	fmt.Println("values of a,b:", a, b)
	*a, *b = *b, *a
	fmt.Println("updated values of a,b:", *a, *b)
}

// Task: Write a function that takes a pointer to an integer, doubles its value, and prints the result.
func DoubleTheValue(r *int) {
	fmt.Println("value before double:", *r)
	*r = *r * 2
	fmt.Println("value after double:", *r)
}

// Task: Demonstrate how pointers can be used to modify elements in an array.
func ModifyArray(arr *[5]int, index int, value int) {
	if index >= 0 && index < len(*arr) { // Ensure index is valid
		(*arr)[index] = value // Modify the array at the specified index
	} else {
		fmt.Println("Index out of range")
	}
}
