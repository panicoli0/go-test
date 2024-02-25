package helper

import "strings"

func UserInputValidation(firstName string, LastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(LastName) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
