package main

import (
	"booking-app/helper"
	"fmt" //means format for Input Output
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // this is how to initialize a List of maps

// simil to class in java
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {
		firstName, LastName, email, userTickets := getUserInputs()

		// user input validations:
		isValidName, isValidEmail, isValidTicketNumber := helper.UserInputValidation(firstName, LastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//slice simil List<T>
			//remainingTickets, bookings = bookTickets(userTickets, firstName, LastName, email)
			bookTickets(userTickets, firstName, LastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, LastName, email)
			//like a foreach
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v \n", firstNames)

			if remainingTickets <= 0 {
				fmt.Printf("We are sold out!")
				break
			}
		} else {
			fmt.Printf("Your input data is wrong!\n")
		}
	}
}

func bookTickets(userTickets uint, firstName string, LastName string, email string) (uint, []UserData) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        LastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v %v booked %v tickets. you will get a confirmation email at %v \n", firstName, LastName, userTickets, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	return remainingTickets, bookings
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var LastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your LastName: ")
	fmt.Scan(&LastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, LastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func greetUser() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have a total of %v remaining and %v are still available\n", conferenceTickets, remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                        // delay of 15sec
	var ticket = fmt.Sprintf("%v tickets for %v, %v", userTickets, firstName, lastName) // x tickets for xUserName xLastName
	fmt.Println("##############")
	fmt.Printf("Sending %v this confirmation will be send to: %v \n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
