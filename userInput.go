// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program takes the users street number and name and shows the users address

package main

import "fmt"

func main() {
	// This function takes the users street number and name and shows the users address
	var streetNumber int
	var streetName string

	// input
	fmt.Println("This program gets a users address.")
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println("Your address is", streetNumber, streetName, ".")
}
