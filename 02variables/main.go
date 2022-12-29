package main

import (
	"fmt"
)

func main() {
	var username string = "hello" // declaring the variable
	fmt.Println(username)
	fmt.Printf("Variable is a type : %T \n", username) //getting the type of varibale  %T used for type

	/*

		NOTE YOU HAVE TO USE THE VARIABLE AFTER DECLARING YOU CANT LEAVE VARIBALE LIKE ANY WHERE
		IT IS A ERROR IN GO LANG

	*/
	var num int = 24
	fmt.Printf("variable is : %T \n", num)
	var logged bool = false
	fmt.Printf("variable is :%T \n", logged)
	var varity uint8 = 243 /// can only store value from 0 to 243
	fmt.Printf("variable is :%T \n", varity)
	var smallFloat float32 = 253.342342342342342
	fmt.Println(smallFloat) //gonna print only decimal places upto 5
	fmt.Printf("varibale is : %T \n", smallFloat)
	var largeFloat float64 = 234.342342342342342
	fmt.Println(largeFloat) //gonna print only decimal places upto 5
	fmt.Printf("varibale is : %T \n", largeFloat)

	var anotherVariable string
	fmt.Println(anotherVariable)

	/*

		by default value of int is 0 {not garbage}
		and by default value of string is {blank or \n}
	*/

	// implicity type

	var website = "amazon.in" // you dont have to declare a type for website
	fmt.Println(website)

	//  no var style

	number_of_user := 300000 // you can declare without var
	fmt.Println(number_of_user)
	/*
		:= THIS OPERATOR CALLED AS WALRUS OPERATOR
		BUT YOU CAN ONLY USE THIS INSIDE THE METHOD LIKE FUNCTION OR SOMETHING LIKE METHODS BUT YOU CANT USE OUTSIDE OF METHOD

		LIKE I CANNOT DECLARE WALRUS OPERATOR OUTSIDE THE MAIN FUNCTION BUT INSIDE THE MAIN FUNCTION I CAN

	*/

	// CAN DECLARE THE CONSTANT ALSO

	const LOgin string = "dfsf"
	fmt.Println(LOgin)
}
