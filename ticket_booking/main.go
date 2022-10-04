package main

import (
	"booking_app/helper"
	"fmt"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// isValidCity := city == "America" || city == "India"

		if isValidName && isValidTicketNumber && isValidEmail {

			bookTicket(userTickets, firstName, lastName, email)

			//call fuction print first name
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next time.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name and last name you enter is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you enter not conatain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number you enter is Invalid")
			}
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome To %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string
	// ask for input
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of Tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
