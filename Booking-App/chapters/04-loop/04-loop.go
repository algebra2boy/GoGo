package main

import (
	"fmt"
	"strings"
)

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
	runLoop()
}

func runLoop() {
	const officeHourName string = "Booking.com"
	var officeHourStart int = 9
	officeHourSession := 8

	// bookings := []string{"Apple", "Banana", 4} // This will not compile since it contains different types
	// bookings := [20]string{} // 20 is the length of the array
	bookings := []string{} // slice, dyanmic array with no limit

	fmt.Printf("Welcome to %v at %vam for %v session \n", officeHourName, officeHourStart, officeHourSession)

	for {

		var firstname, lastname, email, userName string

		fmt.Println("Enter your first name: ")
		fmt.Scanln(&firstname)

		fmt.Println("Enter your last name: ")
		fmt.Scanln(&lastname)

		fmt.Println("Enter your email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter your username: ")
		fmt.Scanln(&userName)
		bookings = append(bookings, firstname+" "+lastname)

		firstnames := []string{}
		for _, booking := range bookings {
			// splits the string with white space as separator
			// for example: ["apple pie"] -> ["apple", "pie"]
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
		fmt.Printf("the first names of booking are: %v\n", firstnames)

		fmt.Printf("Welcome %v %v, your email is %v and your username is %v \n", firstname, lastname, email, userName)
	}
}
