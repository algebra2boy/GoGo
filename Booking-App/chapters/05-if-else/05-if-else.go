package main

import (
	"fmt"
	"strings"
)

func main() {
	const officeHourName string = "Booking.com"
	var officeHourStart int = 9
	officeHourSession := 8

	bookings := []string{} // slice, dyanmic array with no limit

	fmt.Printf("Welcome to %v at %vam for %v session \n", officeHourName, officeHourStart, officeHourSession)

	for {

		var firstname, lastname, email, userName string
		var peopleInFrontQueue int

		fmt.Println("Enter your first name: ")
		fmt.Scanln(&firstname)

		fmt.Println("Enter your last name: ")
		fmt.Scanln(&lastname)

		fmt.Println("Enter your email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter your username: ")
		fmt.Scanln(&userName)

		fmt.Println("Enter the number of people in front of you in the queue: ")
		fmt.Scanln(&peopleInFrontQueue)

		isValidName := len(firstname) >= 2 && len(lastname) >= 2
		isValidEmail := strings.Contains(email, "@")

		if isValidName && isValidEmail {
			bookings = append(bookings, firstname+" "+lastname)

			firstnames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("the first names of booking are: %v\n", firstnames)

			fmt.Printf("Welcome %v %v, your email is %v and your username is %v \n", firstname, lastname, email, userName)
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			continue
		}
	}
}
