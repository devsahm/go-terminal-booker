package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

var bookings []string

func main() {

	// bookings := []string{}

	fmt.Printf("conference name is %T, Conference Tickets is %T, Conference Tickets Remaining is %T\n", conferenceName, conferenceTickets, remainingTickets)

	//Greet Users
	greetUser()

	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTicket(userTicket, firstName, lastName, email)

			//Print the firstname of the user that booked
			firstNames := firstNames()

			fmt.Printf("The first name of the bookings is: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the  loop
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("The first name or lastname must not be less than 2 character")
			}

			if !isValidEmail {
				fmt.Println("The email address does not contain the @  sign")
			}

			if !isValidTickets {
				fmt.Println("The user tickets entry is not valid")
			}

		}
	}

}

func greetUser() {

	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func firstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var email string
	var lastName string
	var userTicket uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of Tickets")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTicket

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for entering  %v tickets. You will recieve a confirmation email into %v for your request\n", firstName, lastName, userTicket, email)

	fmt.Printf("%v tickets is remaining. Thank you for booking in %v\n", remainingTickets, conferenceName)

}
