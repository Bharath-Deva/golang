package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Actual Value of pointer pointing to is ", *ptr)

	// ptr = &(*ptr + 2) // error & cannot get the address of a operation. It can get the address of the variable
	*ptr = *ptr + 2 // * will get the value from the address stroed in the pointer. Similarly when it is placed before '=' it will store the value of the operation in the address stored in the pointer
	// So basically * will decipher the pointer into its memory location based on where it is placed on the '=' sign

	fmt.Println("New value is: ", myNumber)

}
