package main

import "fmt"

func main() {
	fmt.Println("maps in go")

	languages := make(map[string]string)
	/*
		here we use make keyword to init map in go
		so how we define these lets have a look

		make(map[string]string)

		make : keyword
		map :written to define that we are goimg to use map
		[string ] :we are going to use key of string type
		string {note with no brackets} = we are going to use string for values
	*/

	/*

		NOW LET START ADDING NEW KEY AND VALUE PAIR
	*/

	languages["js"] = "javascript lover"
	languages["go"] = "we are learning go lang"

	/*
		so you get it what we are doing
	*/
	fmt.Println(languages)

	/*

		now how to delete value form go lang
		we use delete keyword
	*/
	delete(languages, "js")
	fmt.Println(languages)
}
