package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	// "example.com/goodJob"
	// "example.com/thumb"
)

func main() {
	/*
		Set properties of the predefined Logger,
		including the log entry prefix and a flag to disable printing
		the time, source file, and line number.
	*/

	// Request a greeting message.
	message, err := greetings.Hello("Eastwoodli")

	// If an error was returned, print it to the console and exit the program

	if err != nil {
		log.SetPrefix("greetings.Hello:")
		log.SetFlags(0)
		log.Fatal(err)
	}
	
	// Get a greeting message and print it.
	// message:=greetings.Hello("Eastwoodli")
	// message:= thumb.Thumb("Eastwoodli")
	// fmt.Println(message)

	// If no error was returned, print the returned message to the console
	fmt.Println(message)

}
