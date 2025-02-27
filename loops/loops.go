package loops

import "fmt"

func Oddeven(arr1 [9]int) {
	for i := range len(arr1) {
		if arr1[i]%2 == 0 {
			fmt.Println(arr1[i], " is even")
		} else {
			fmt.Println(arr1[i], "is odd")
		}

	}

}

func EvenOdd(arr1 [9]int) {
	for i := range len(arr1) {
		if arr1[i]%2 == 0 {
			fmt.Println(arr1[i], " is even")
		} else {
			fmt.Println(arr1[i], "is odd")
		}

	}

}

// Oddeven accepts a slice of integers and checks if each element is even or odd
func SliceOddeven(arr2 []int) {
	for i := range arr2 {
		if arr2[i]%2 == 0 {
			fmt.Println(arr2[i], "is even")
		} else {
			fmt.Println(arr2[i], "is odd")
		}
	}
}

//Task: Write a program that prints the multiplication table (1–10) for a number entered by the user.
func TableMultiplication(a int) {
	for i := 1; i <= 10; i++ {
		//fmt.Println(a, "*", i, "=", a*i) Use below for best coding standards
		fmt.Printf("%d * %d = %d\n", a, i, a*i)
	}
}

//Task: Iterate over an array of integers and calculate the sum of all elements using for loops.
func SumOfElements(arr3 []int) int {
	sum := 0
	/*for i := 0; i < len(arr3); i++ {
		sum += arr3[i]
	}
	fmt.Println("sum of arr3:", sum)*/
	for _, value := range arr3 {
		sum += value
	}
	return sum
}

//Task: Use a for loop to find all prime numbers between 1 and 100.
func PrimeNums() {
	primes := []int{} // Create an empty slice to store prime numbers

	// Loop through numbers from 2 to 100
	for n := 2; n <= 100; n++ {
		isPrime := true // Assume the number is prime

		// Check if `n` is divisible by any number from 2 to √n
		for i := 2; i*i <= n; i++ {
			if n%i == 0 { // If `n` is divisible, it's not prime
				isPrime = false
				break
			}
		}

		// If the number is prime, add it to the slice
		if isPrime {
			primes = append(primes, n)
		}
	}
	fmt.Println("Prime numbers from 1 to 100:", primes)
}

//Approach 2 best approach
/*func PrimeNums(limit int) {
	if limit < 2 {
		fmt.Println("No prime numbers in this range.")
		return
	}

	primes := []int{2} // Start with 2 as it's the first prime

	// Loop through only ODD numbers after 2 (no need to check even numbers)
	for n := 3; n <= limit; n += 2 {
		isPrime := true

		// Check divisibility only up to √n
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		// If prime, add to slice
		if isPrime {
			primes = append(primes, n)
		}
	}

	// Print prime numbers
	fmt.Println("Prime numbers up to", limit, ":", primes)
}

func main() {
	PrimeNums(100) // Call function with limit 100
}*/
