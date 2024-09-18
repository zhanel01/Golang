package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

}

func variables() {
	var a int = 10
	var b float64 = 20.5
	var c string = "Hello"
	var d bool = true

	e := 5 // short declaration syntax
	fmt.Printf("a: %d, type: %T\n", a, a)
	fmt.Printf("b: %.2f, type: %T\n", b, b)
	fmt.Printf("c: %s, type: %T\n", c, c)
	fmt.Printf("d: %t, type: %T\n", d, d)
	fmt.Printf("e: %d, type: %T\n", e, e)
}

// ex3
func add(a int, b int) int {
	return a + b
}

func swap(a string, b string) (string, string) {
	return b, a
}

func divide(x int, y int) (int, int) {
	return x / y, x % y
}

func functions() {
	// Using the add function
	sum := add(3, 4)
	fmt.Println("Sum:", sum)

	// Using the swap function
	a, b := swap("Hello", "World")
	fmt.Println("Swapped:", a, b)

	// Using the divide function
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)
}

//ex4
func controlStructures() {

	// Check if number is positive, negative, or zero
	var num int
	fmt.Println("Enter number:")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// Calculate the sum of the first 10 natural numbers
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of the first 10 natural numbers:", sum)

	// Switch statement for the day of the week
	var day int
	fmt.Println("Enter a number (1-7):")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input")
	}
}
