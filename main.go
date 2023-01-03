package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"

	const conferenceTickets int = 50

	var remainingTickets uint = 50

	// var bookings []string

	bookings := []string{}

	fmt.Printf("conference name is %T, Conference Tickets is %T, Conference Tickets Remaining is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		isValiEmail := strings.Contains(email, "@")

		isValidTickets := userTicket > 0 && userTicket <= remainingTickets

		if isValiEmail && isValidName && isValidTickets {

			remainingTickets = remainingTickets - userTicket

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for entering  %v tickets. You will recieve a confirmation email into %v for your request\n", firstName, lastName, userTicket, email)

			fmt.Printf("%v tickets is remaining. Thank you for booking in %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {

				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])
			}

			if remainingTickets == 0 {
				//end the  loop
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}

			fmt.Printf("The first name of the bookings is: %v\n", firstNames)

		} else {

			if !isValidName {
				fmt.Println("The first name or lastname must not be less than 2 character")
			}

			if !isValiEmail {
				fmt.Println("The email address does not contain the @  sign")
			}

			if !isValidTickets {
				fmt.Println("The user tickets entry is not valid")
			}

		}
	}

}
