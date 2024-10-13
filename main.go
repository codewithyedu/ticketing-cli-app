package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var (
	remainingTickets = conferenceTickets
	conferenceName   = "Go conference"
	wg               = sync.WaitGroup{}
	bookings         []UserData
)

// UserData struct to store information about each booking
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)

			// Set the number of goroutines to wait for
			wg.Add(1)
			// Launch a goroutine to send the ticket confirmation email
			go sendTicket(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("These are our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sold out")
				break
			}

		} else {
			// Display error messages for invalid inputs
			if !isValidName {
				fmt.Println("Error: Name too short.")
			}

			if !isValidEmail {
				fmt.Println("Error: Invalid email.")
			}

			if !isValidTicketNumber {
				fmt.Printf("Error: Invalid ticket number. Tickets remaining: %v", remainingTickets)
			}

			fmt.Println("Please try again.")
		}

	}

	// Wait for all goroutines to finish before exiting the program
	wg.Wait()

}

func greetUsers() {
	// for getting the type use %T -> fmt.Printf("Type of conferenceName is %T", conferenceName) -> string
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf(
		"We have a total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets,
	)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	_, err1 := fmt.Scan(&firstName)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("Enter your last name:")
	_, err2 := fmt.Scan(&lastName)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("Enter your email address:")
	_, err3 := fmt.Scan(&email)
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("Enter number of tickets:")
	_, err4 := fmt.Scan(&userTickets)
	if err4 != nil {
		fmt.Println(err4)
	}

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	remainingTickets -= userTickets

	fmt.Printf(
		"Thank you %v %v for purchasing %v tickets. A confirmation email will be sent to %v.\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		firstName := booking.firstName
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("##########")

	// Decrement the WaitGroup counter to indicate this goroutine is complete
	wg.Done()
}
