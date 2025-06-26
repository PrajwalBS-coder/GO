package main

import (
	"fmt"
	"strings"
)

func main() {
	var confna = "Go Conferenece"
	const tickets int = 50
	var remaintickets uint = 50
	var fname, lname, email string
	var utickets uint
	bookings := []string{}

	fmt.Println("GO Project!")
	fmt.Printf("Welcome to %s!\n", confna)
	fmt.Printf("Get Your Tickets There are %d Tickets Availabel\n", remaintickets)

	//taking inputs from user
	for {
		fmt.Println("Enter Your Name")
		fmt.Scan(&fname)
		fmt.Println("Enter Your Name")
		fmt.Scan(&lname)
		fmt.Println("Enter Your Email")
		fmt.Scan(&email)
		fmt.Println("Enter the no.of tickets")
		fmt.Scan(&utickets)
		var temp= remaintickets - utickets
		remaintickets =temp
		bookings = append(bookings, fname+" "+lname)
		fmt.Printf("Thank You %v %v for booking %v no.of tickets you willl confirmation email to this %v email ", fname, lname, utickets, email)
		fmt.Printf("\n%v Tickets remaining for %v", remaintickets, confna)

		fn := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			fn = append(fn, names[0])

		}

		fmt.Printf("\nAll the Bookings thart are Done %v\n", fn)

	}

}
