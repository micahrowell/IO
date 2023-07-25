package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//userFirstNameMessage := "Please type your first name below:"
	//fmt.Println(userFirstNameMessage)
	//
	//var firstName string
	//fmt.Scanln(&firstName)
	//
	//userLastNameMessage := "Please type your last name below:"
	//fmt.Println(userLastNameMessage)
	//
	//var lastName string
	//fmt.Scanln(&lastName)
	//
	//fmt.Println("Your first name is:")
	//fmt.Println(firstName)
	//fmt.Println("Your last name is:")
	//fmt.Println(lastName)

	// I want to get the user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name below:")

	userName, _ := reader.ReadString('\n')

	fmt.Println("The name that you entered is:")
	fmt.Println(userName)
}
