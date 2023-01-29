package main

import "fmt"

func main() {
	fmt.Println("if and else in golang")
	var msg string
	loginCount := 1
	if loginCount < 10 {
		msg = "haav a day"
	} else if loginCount > 10 {
		msg = "watch out baby"
	} else {
		msg = "hehehehhehe"
	}

	/*

		you cnat write if in go like this
		if loginCount < 10
		{
		msg = "haav a day"
		}
		it is an error in the golang


	*/
	fmt.Println(msg)
	if num := 3; num == 3 {
		fmt.Println("you are correct")
	}
	/*
		you can also use num like this
		in this you defined the num with 3 here first and
		then you get back to check the condition
	*/

}
