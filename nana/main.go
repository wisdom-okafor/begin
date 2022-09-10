package main

import "fmt" 

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50 
	var remainingTickets uint = 50 

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	// scan(&userName)

	fmt.Println(remainingTickets)
    fmt.Println(&remainingTickets)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	
}	