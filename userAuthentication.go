package main

import (
	"fmt"
	"strings"
)

var message, user string

func main() {
	//Insert Code here

	fmt.Printf("Input Username\n")
	fmt.Scan(&user)

	if strings.ToLower(user) == "admin" {
		message = "Welcome, Admin!"
	} else if strings.ToLower(user) == "robin" || strings.ToLower(user) == "john" {
		message = "Welcome! " + user
	} else {
		message = "Merry Men"
	}
	println(message)
}
