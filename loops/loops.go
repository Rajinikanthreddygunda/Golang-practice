package loops

import "fmt"

func Oddeven(arr1 [9]int) {
	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 == 0 {
			fmt.Println(arr1[i], " is even")
		} else {
			fmt.Println(arr1[i], "is odd")
		}

	}

}

// Oddeven accepts a slice of integers and checks if each element is even or odd
func SliceOddeven(arr2 []int) {
	for i := 0; i < len(arr2); i++ {
		if arr2[i]%2 == 0 {
			fmt.Println(arr2[i], "is even")
		} else {
			fmt.Println(arr2[i], "is odd")
		}
	}
}
