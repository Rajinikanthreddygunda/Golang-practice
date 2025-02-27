// Assignment 2: Temperature Converter
// Use these: Methods, Receivers, Pointers
// Problem Statement:Create a program that converts temperatures between Celsius and Fahrenheit.
// Requirements:
// 1. Define a struct Temperature with a Value (float64).
// 2. Implement methods:
//   - ToFahrenheit() → Converts from Celsius to Fahrenheit.
//   - ToCelsius() → Converts from Fahrenheit to Celsius.
//
// 3. Use pointers to modify the temperature directly.
// Formula:
// * °F = (°C × 9/5) + 32
// * °C = (°F - 32) × 5/9
package main

import "fmt"

type Temperature struct {
	value float64
}

func (t *Temperature) ToFahrenheit() {
	t.value = (t.value * 9 / 5) + 32
}

func (t *Temperature) ToCelsius() {
	t.value = (t.value - 32) * 5 / 9
}

func main() {
	var tempValue float64
	var choice int

	fmt.Println("Enter the temperature value:")
	fmt.Scan(&tempValue)

	// Create a Temperature instance with the user's input
	temp := &Temperature{value: tempValue}

	// Ask user for conversion choice
	fmt.Println("Choose an option:")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Fahrenheit to Celsius")
	fmt.Scan(&choice)

	// Perform conversion based on user's choice
	switch choice {
	case 1:
		temp.ToFahrenheit()
		fmt.Printf("Converted Temperature: %.2f°F\n", temp.value)
	case 2:
		temp.ToCelsius()
		fmt.Printf("Converted Temperature: %.2f°C\n", temp.value)
	default:
		fmt.Println("Invalid choice. Please select either 1 or 2.")
	}

}
