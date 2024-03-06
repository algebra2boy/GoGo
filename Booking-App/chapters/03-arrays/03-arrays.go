package main

import "fmt"

func main() {

	const officeHourName string = "Booking.com"
	var officeHourStart int = 9
	officeHourSession := 8

	// bookings := []string{"Apple", "Banana", 4} // This will not compile since it contains different types
	// bookings := [20]string{} // 20 is the length of the array
	bookings := []string{} // slice, dyanmic array with no limit

	fmt.Printf("Welcome to %v at %vam for %v session \n", officeHourName, officeHourStart, officeHourSession)

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

	fmt.Printf("Current bookings: %v\n", bookings)
	fmt.Printf("first booking: %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings) // []string
	fmt.Printf("Array length %v\n", len(bookings))

	fmt.Printf("Welcome %v %v, your email is %v and your username is %v \n", firstname, lastname, email, userName)
}
