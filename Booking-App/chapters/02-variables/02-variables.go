package main

import "fmt"

func main() {

	const officeHourName string = "Booking.com"
	var officeHourStart int = 9
	officeHourSession := 8

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

	fmt.Printf("Welcome %v %v, your email is %v and your username is %v \n", firstname, lastname, email, userName)
}
