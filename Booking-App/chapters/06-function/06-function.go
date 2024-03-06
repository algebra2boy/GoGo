package main

import (
	"fmt"
	"strings"
)

const officeHourName string = "Booking.com"

var officeHourStart int = 9
var bookings = []string{} // slice, dyanmic array with no limit

func main() {

	greetUser()

	for {

		firstname, lastname, email, peopleInFrontQueue := getUserInput()

		isValidName, isValidEmail, isValidNumber := validateUserInput(firstname, lastname, email, peopleInFrontQueue)

		if isValidName && isValidEmail && isValidNumber {
			bookings = append(bookings, firstname+" "+lastname)

			firstnames := printFirstNames()
			fmt.Printf("the first names of booking are: %v\n", firstnames)

			fmt.Printf("Welcome %v %v, your email is %v\n", firstname, lastname, email)
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidNumber {
				fmt.Println("Invalid number")
			}
			continue
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v at %vam \n", officeHourName, officeHourStart)
}

func getUserInput() (string, string, string, int) {
	var firstname, lastname, email string
	var peopleInFrontQueue int

	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastname)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter the number of people in front of you in the queue: ")
	fmt.Scanln(&peopleInFrontQueue)

	return firstname, lastname, email, peopleInFrontQueue
}

func validateUserInput(firstName string, lastName string, email string, peopleInFrontQueue int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := peopleInFrontQueue >= 0

	return isValidName, isValidEmail, isValidNumber
}

func printFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}

	return firstnames
}
