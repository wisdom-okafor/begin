package main

import ("fmt"
        "strings"
)		 


func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50 
	var remainingTickets uint = 50 
	var bookings = []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email sddress:")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
	
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Thes first names of bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println(("Our conference is blocked out. Come back next year"))
			break
		}
	}
}