package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValiEmail := strings.Contains(email, "@")

	isValidTickets := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValiEmail, isValidTickets

}
