package main

import "strings"

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
) (isValidName, isValidEmail, isValidTicketNumber bool) {
	// validateUserInput checks if the user input is valid.
	// It returns three boolean values indicating the validity of the name, email, and ticket number.
	isValidName = (len(firstName) >= 2) && (len(lastName) >= 1)
	isValidEmail = strings.Contains(email, "@")
	isValidTicketNumber = (userTickets > 0) && (userTickets <= remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}
