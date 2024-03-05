package main

import "fmt"

// In Go, function names starting with an uppercase letter are exported (visible outside the package):

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
