package main

import "fmt"

// Define the Employee struct
type Employee struct {
	Name string
	ID   int
}

// Method for the Employee struct
func (e Employee) Work() {
	fmt.Printf("Employee Name: %s, ID: %d\n", e.Name, e.ID)
}

// Define the Manager struct, embedding Employee
type Manager struct {
	Employee
	Department string
}

func main() {
	manager := Manager{
		Employee:   Employee{Name: "Bob", ID: 101},
		Department: "Sales",
	}
	manager.Work()
}
