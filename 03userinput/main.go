package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "we are going to take input from user" /// VAR STORE USING WALRUS OPERATOR
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) /* /// we are using bufio libraray to read the buffer which we eventually get from the user
	and to read from the user we will use os library also
	Stdin means standard input
	*/
	fmt.Println("enter the reading for the pizaa:")

	// comma ok // err ok

	input, _ := reader.ReadString('\n') /// using just like javascript
	/*

		WE DECLARE INPUT AND _  BUT WHY WE NEED TWO VARIABLE(input and _)
		BECAUSE ReadString can give two possible output one is result from user and other gonna be error so the UNDERSOCRE(_) is responsible to hold the error and this is call comma err okay syntax


	*/
	fmt.Println(input)
	fmt.Printf("TYPE OF The STRING : %T \n", input)
	/// the value from user end up being string
	/*
		input from user always gonna return in the format of string
		so if we want a number how we gonna do that

		so we will use conversion for that you can check in next moddule
	*/
}
