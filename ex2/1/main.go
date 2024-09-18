package main

import "fmt"

// Define the Person struct
type Person struct {
	Name string
	Age  int
}

// Method for the Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	person.Greet()
}
