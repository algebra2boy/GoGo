package main

import (
	"fmt"
)

// this is a comment
func main() {
	fmt.Println("Hello, World!")

	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	var x string = "Hello, World!" // explicit type
	var y = "Hello, World!"        // infer type
	z := "Hello, World!"           // infer shorthand type
	fmt.Println(x, y, z)
	fmt.Println(x == y) // string comparison
	fmt.Printf("how are you %v\n", x)

	const a string = "Hello, World!" // constant (cannot be changed)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
	fmt.Println(a)
}
