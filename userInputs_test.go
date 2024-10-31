package main

import "testing"

func TestInvalidEmail(t *testing.T) {
	var (
		fName            = "Yedukrishna"
		lName            = "R"
		email            = "some email"
		userTickets uint = 12
	)

	_, isValidEmail, _ := validateUserInput(fName, lName, email, userTickets)
	if isValidEmail {
		t.Errorf("TestInvalidEmail Failed: '%v' is not a valid email.", email)
	}
}

func TestInvalidName(t *testing.T) {
	var (
		fName            = "Y"
		lName            = "R"
		email            = "cwy@email.com"
		userTickets uint = 12
	)

	isValidName, _, _ := validateUserInput(fName, lName, email, userTickets)
	if isValidName {
		t.Errorf("TestInvalidName Failed: '%v %v' is not a valid name.", fName, lName)
	}
}

func TestInvalidUserTickets(t *testing.T) {
	var (
		fName            = "Yedukrishna"
		lName            = "R"
		email            = "cwy@email.com"
		userTickets uint = 0
	)

	_, _, isValidTicketNumber := validateUserInput(fName, lName, email, userTickets)
	if isValidTicketNumber {
		t.Errorf("TestInvalidTicketNumber Failed: '%v' is not a valid ticket number.", userTickets)
	}
}
