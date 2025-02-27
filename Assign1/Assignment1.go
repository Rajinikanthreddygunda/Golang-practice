// Assignment 1: Employee Management System
// Use these: Structs, Methods, Pointers, Slices, Maps
// Problem Statement: Create a simple Employee Management System that allows adding employees, updating their salary, and displaying their details.
// Requirements:
// 1. Define a struct Employee with fields:
//     * ID (int)
//     * Name (string)
//     * Salary (float64)
// 2. Implement methods:
//     * UpdateSalary(newSalary float64) → Updates the salary of an employee.
//     * Display() → Prints employee details.
// 3. Use a map to store employees (map[int]*Employee).
// 4. Implement functions:
//     * AddEmployee(emp *Employee, db map[int]*Employee) → Adds an employee to the map.
//     * RemoveEmployee(id int, db map[int]*Employee) → Removes an employee by ID.

package main

import "fmt"

// Employee struct to store employee details
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Method to update the salary of an employee
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

// Method to display employee details
func (e Employee) Display() {
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

// Function to add an employee to the map
func AddEmployee(emp *Employee, db map[int]*Employee) {
	if _, exists := db[emp.ID]; exists {
		fmt.Println("Employee with ID already exists.")
		return
	}
	db[emp.ID] = emp
	fmt.Println("Employee added successfully!")
}

// Function to remove an employee from the map
func RemoveEmployee(id int, db map[int]*Employee) {
	if _, exists := db[id]; !exists {
		fmt.Println("Employee with this ID does not exist.")
		return
	}
	delete(db, id)
	fmt.Println("Employee removed successfully!")
}

func main() {
	// Creating a map to store employees
	employees := make(map[int]*Employee)

	// Adding employees
	emp1 := &Employee{ID: 101, Name: "Alice", Salary: 50000}
	emp2 := &Employee{ID: 102, Name: "Bob", Salary: 60000}

	AddEmployee(emp1, employees)
	AddEmployee(emp2, employees)

	// Display all employees
	fmt.Println("\nEmployee List:")
	for _, emp := range employees {
		emp.Display()
	}

	// Updating salary of an employee
	fmt.Println("\nUpdating Salary of Employee ID 101:")
	employees[101].UpdateSalary(55000)
	employees[101].Display()

	// Removing an employee
	fmt.Println("\nRemoving Employee with ID 102:")
	RemoveEmployee(102, employees)

	// Displaying updated employee list
	fmt.Println("\nUpdated Employee List:")
	for _, emp := range employees {
		emp.Display()
	}
}
