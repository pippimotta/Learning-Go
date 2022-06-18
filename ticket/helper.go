package main

import "strings"

func ValidateUserInput(fn string, ln string, email string, ut uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fn) >= 2 && len(ln) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := ut > 0 && ut <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
